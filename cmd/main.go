package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
	chiDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/tomsiewert/elastiflow-to-netmeta/internal/misc"
)

var (
	generateSQL         = flag.Bool("generate-sql", true, "Generate SQL from OpenSearch Data")
	dryRun              = flag.Bool("dry-run", false, "Dry Run")
	connectToClickhouse = flag.Bool("connect-to-clickhouse", false, "Connect to Clickhouse directly instead of generating large SQL dumps")
	clickhouseHost      = flag.String("clickhouse-host", "127.0.0.1:9000", "Clickhouse Host (required if connecto-to-clickhouse is enabled)")
	clickhouseDB        = flag.String("clickhouse-db", "default", "")
	clickhouseUser      = flag.String("clickhouse-user", "default", "")
	clickhousePass      = flag.String("clickhouse-password", "", "")
	batchSize           = flag.Int("batch-size", 1000, "Number of documents processed per Batch")
	workerCount         = flag.Int("batch-workers", 10, "Number of workers to process batches")

	enableLog = flag.Bool("enable-log", false, "")
	year      = flag.Int("year", 2023, "")
	month     = flag.Int("month", 1, "")
	startDay  = flag.Int("start-day", 1, "")
	endDay    = flag.Int("end-day", 31, "")
)

func connect() (chiDriver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{*clickhouseHost},
			Auth: clickhouse.Auth{
				Database: *clickhouseDB,
				Username: *clickhouseUser,
				Password: *clickhousePass,
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "elastiflow-to-netmeta", Version: "0.1"},
				},
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}

func processLines(lines []string, conn chiDriver.Conn, ctx *context.Context) {
	for _, line := range lines {
		var hit OSFlowResult
		err := json.Unmarshal([]byte(line), &hit)
		if err != nil {
			log.Println(err)
			continue
		}

		if *enableLog {
			log.Println(hit.Fields.Flow_Collect_Timestamp[0])
		}
		// Default values
		flowType := "FLOWUNKNOWN"
		flowDirecction := "0"
		packets := "1"
		eType := "0"
		srcMac := "0"
		dstMac := "0"
		ingressVrfId := "0"
		egressVrfId := "0"
		ipTos := "0"
		forwardingStatus := "0"
		icmpType := "0"
		icmpCode := "0"
		ipv6FlowLabel := "0"
		fragmentId := "0"
		fragmentOffset := "0"
		biFlowDirection := "0"
		nextHop := "::"
		nextHopAS := "0"
		srcNet := "0"
		dstNet := "0"

		protoSave, _ := ProtocolMap[hit.Source.L4_Proto_Name]
		// Dynamic values
		date := fmt.Sprint(hit.Fields.Flow_Collect_Timestamp[0].Unix()) // Date, TimeReceived, TimeFlowStart, TimeFlowEnd
		sequenceNum := fmt.Sprint(hit.Source.Sflow_sample_seqNum)
		samplingRate := fmt.Sprint(hit.Source.Flow_Packets)
		bytes := fmt.Sprint(hit.Source.Ip_Packet_Size)
		proto := fmt.Sprint(protoSave)
		srcPort := fmt.Sprint(hit.Source.Flow_Src_L4_Port_ID)
		dstPort := fmt.Sprint(hit.Source.Flow_Dst_L4_Port_ID)
		inIf := fmt.Sprint(hit.Source.Flow_In_Netif_Index)
		outIf := fmt.Sprint(hit.Source.Flow_Out_Netif_Index)
		srcVlan := fmt.Sprint(hit.Source.Vlan_cTag_id)
		dstVlan := fmt.Sprint(hit.Source.Vlan_cTag_id)
		vlanId := fmt.Sprint(hit.Source.Vlan_cTag_id)
		ipTTL := fmt.Sprint(hit.Source.Ip_Ttl)
		tcpFlags := fmt.Sprint(hit.Source.Tcp_Flags_Bits)
		srcAS := fmt.Sprint(hit.Source.Flow_Src_As_Asn)
		dstAS := fmt.Sprint(hit.Source.Flow_Dst_As_Asn)
		samplerAddress := misc.ConvertIPv4ToIPv6(hit.Source.Flow_Export_Ip_Addr)
		srcAddr := misc.ConvertIPv4ToIPv6(hit.Source.Flow_Src_Ip_Addr)
		dstAddr := misc.ConvertIPv4ToIPv6(hit.Source.Flow_Dst_Ip_Addr)

		query := "INSERT INTO `flows_raw` (`Date`, `FlowType`, `SequenceNum`, `TimeReceived`, `SamplingRate`, `FlowDirection`, `SamplerAddress`, `TimeFlowStart`, `TimeFlowEnd`, `Bytes`, `Packets`, `SrcAddr`, `DstAddr`, `EType`, `Proto`, `SrcPort`, `DstPort`, `InIf`, `OutIf`, `SrcMac`, `DstMac`, `SrcVlan`, `DstVlan`, `VlanId`, `IngressVrfId`, `EgressVrfId`, `IPTos`, `ForwardingStatus`, `IPTTL`, `TCPFlags`, `IcmpType`, `IcmpCode`, `IPv6FlowLabel`, `FragmentId`, `FragmentOffset`, `BiFlowDirection`, `SrcAS`, `DstAS`, `NextHop`, `NextHopAS`, `SrcNet`, `DstNet`) VALUES (toDateTime('%s'), '%s', %s, %s, %s, %s, '%s', %s, %s, %s, %s, '%s', '%s', %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, '%s', %s, %s, %s);"

		finalQuery := fmt.Sprintf(query,
			date, flowType, sequenceNum, date, samplingRate,
			flowDirecction, samplerAddress, date, date,
			bytes, packets, srcAddr, dstAddr, eType, proto, srcPort,
			dstPort, inIf, outIf, srcMac, dstMac, srcVlan, dstVlan,
			vlanId, ingressVrfId, egressVrfId, ipTos, forwardingStatus,
			ipTTL, tcpFlags, icmpType, icmpCode, ipv6FlowLabel,
			fragmentId, fragmentOffset, biFlowDirection, srcAS, dstAS,
			nextHop, nextHopAS, srcNet, dstNet)

		if !*dryRun {
			if *connectToClickhouse {
				if err := conn.AsyncInsert(*ctx, finalQuery, false); err != nil {
					log.Println(err)
				}
			}
		} else {
			if *enableLog {
				fmt.Println(finalQuery)
			}
		}
	}
}

func main() {
	log.Println("Hi Mom")
	flag.Parse()

	days := misc.MakeRange(*startDay, *endDay)
	if *generateSQL {
		var (
			conn chiDriver.Conn
			ctx  = context.Background()
		)

		log.Println("Generate SQL from pre-queried OS-data")

		if !*dryRun && *connectToClickhouse {
			log.Println("Connect to Clickhouse using native driver")
			var err error
			conn, err = connect()
			if err != nil {
				panic((err))
			}
		}
		for _, day := range days {
			dayString := fmt.Sprintf("%04d-%02d-%02d", *year, *month, day)
			log.Println("Inject SQL data to", dayString)

			jsonFile, err := os.Open(dayString + ".json")
			if err != nil {
				log.Fatalln(err)
			}
			defer jsonFile.Close()

			linesChan := make(chan []string, *workerCount)
			var wg sync.WaitGroup

			for i := 0; i < *workerCount; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for lines := range linesChan {
						processLines(lines, conn, &ctx)
					}
				}()
			}

			jsonScanner := bufio.NewScanner(jsonFile)
			var lines []string

			for jsonScanner.Scan() {
				line := jsonScanner.Text()
				lines = append(lines, line)

				if len(lines) == *batchSize {
					linesCopy := append([]string(nil), lines...)
					wg.Add(1)
					linesChan <- linesCopy

					lines = nil
				}
			}

			if err := jsonScanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
				close(linesChan)
				return
			}

			if len(lines) > 0 {
				wg.Add(1)
				linesCopy := append([]string(nil), lines...)
				linesChan <- linesCopy
			}

			close(linesChan)
			wg.Wait()
		}
		if !*dryRun && *connectToClickhouse {
			conn.Close()
		}
	}
}
