package main

import "time"

type OSFlowResult struct {
	Source struct {
		Timestamp                                int      `json:"@timestamp"`
		As_Label                                 []string `json:"as.label"`
		Bgp_Msg_Type_Name                        string   `json:"bgp.msg.type.name,omitempty"`
		Flow_Bytes                               int64    `json:"flow.bytes"`
		Flow_Client_As_Asn                       int      `json:"flow.client.as.asn"`
		Flow_Client_As_Label                     string   `json:"flow.client.as.label"`
		Flow_Client_As_Org                       string   `json:"flow.client.as.org"`
		Flow_Client_Host_Name                    string   `json:"flow.client.host.name"`
		Flow_Client_Ip_Addr                      string   `json:"flow.client.ip.addr"`
		Flow_client_ip_subnet_maskSize           int      `json:"flow.client.ip.subnet.mask_size,omitempty"`
		Flow_Client_L4_Port_ID                   int      `json:"flow.client.l4.port.id,omitempty"`
		Flow_Client_L4_Port_Name                 string   `json:"flow.client.l4.port.name,omitempty"`
		Flow_Client_Mac_Addr                     string   `json:"flow.client.mac.addr,omitempty"`
		Flow_Collect_Timestamp                   int      `json:"flow.collect.timestamp"`
		Flow_Community_ID                        string   `json:"flow.community.id"`
		Flow_Conversation_ID                     string   `json:"flow.conversation.id,omitempty"`
		Flow_Direction_Name                      string   `json:"flow.direction.name"`
		Flow_Dst_As_Asn                          int      `json:"flow.dst.as.asn"`
		Flow_Dst_As_Label                        string   `json:"flow.dst.as.label"`
		Flow_Dst_As_Org                          string   `json:"flow.dst.as.org"`
		Flow_Dst_Host_Name                       string   `json:"flow.dst.host.name"`
		Flow_Dst_Ip_Addr                         string   `json:"flow.dst.ip.addr"`
		Flow_dst_ip_subnet_maskSize              int      `json:"flow.dst.ip.subnet.mask_size,omitempty"`
		Flow_Dst_L4_Port_ID                      int      `json:"flow.dst.l4.port.id,omitempty"`
		Flow_Dst_L4_Port_Name                    string   `json:"flow.dst.l4.port.name,omitempty"`
		Flow_Dst_Mac_Addr                        string   `json:"flow.dst.mac.addr,omitempty"`
		Flow_Export_Host_Name                    string   `json:"flow.export.host.name"`
		Flow_Export_Ip_Addr                      string   `json:"flow.export.ip.addr"`
		Flow_Export_L4_Port_ID                   int      `json:"flow.export.l4.port.id"`
		Flow_Export_Sysuptime                    int      `json:"flow.export.sysuptime"`
		Flow_Export_Type                         string   `json:"flow.export.type"`
		Flow_Export_Version_Name                 string   `json:"flow.export.version.name"`
		Flow_Export_Version_Ver                  int      `json:"flow.export.version.ver"`
		Flow_In_Bytes                            int      `json:"flow.in.bytes"`
		Flow_In_Netif_Index                      int      `json:"flow.in.netif.index"`
		Flow_In_Netif_Name                       string   `json:"flow.in.netif.name"`
		Flow_In_Packets                          int      `json:"flow.in.packets"`
		Flow_In_Vlan_Tag_ID                      int      `json:"flow.in.vlan.tag.id"`
		Flow_In_Vlan_Tag_Pcp_Name                string   `json:"flow.in.vlan.tag.pcp.name"`
		Flow_IsServer                            string   `json:"flow.isServer"`
		Flow_Locality                            string   `json:"flow.locality"`
		Flow_meter_packetSelect_interval_packets int      `json:"flow.meter.packet_select.interval.packets"`
		Flow_meter_packetsDrop                   int      `json:"flow.meter.packets_drop"`
		Flow_meter_packetsTotal                  int      `json:"flow.meter.packets_total"`
		Flow_nextHop_as_asn                      int      `json:"flow.next_hop.as.asn,omitempty"`
		Flow_nextHop_as_label                    string   `json:"flow.next_hop.as.label,omitempty"`
		Flow_nextHop_as_org                      string   `json:"flow.next_hop.as.org,omitempty"`
		Flow_nextHop_host_name                   string   `json:"flow.next_hop.host.name,omitempty"`
		Flow_nextHop_ip_addr                     string   `json:"flow.next_hop.ip.addr,omitempty"`
		Flow_Out_Netif_Index                     int      `json:"flow.out.netif.index,omitempty"`
		Flow_Out_Netif_Name                      string   `json:"flow.out.netif.name,omitempty"`
		Flow_Out_Vlan_Tag_ID                     int      `json:"flow.out.vlan.tag.id"`
		Flow_Out_Vlan_Tag_Pcp_Name               string   `json:"flow.out.vlan.tag.pcp.name"`
		Flow_Packets                             int      `json:"flow.packets"`
		Flow_seqNum                              int      `json:"flow.seq_num"`
		Flow_Server_As_Asn                       int      `json:"flow.server.as.asn"`
		Flow_Server_As_Label                     string   `json:"flow.server.as.label"`
		Flow_Server_As_Org                       string   `json:"flow.server.as.org"`
		Flow_Server_Host_Name                    string   `json:"flow.server.host.name"`
		Flow_Server_Ip_Addr                      string   `json:"flow.server.ip.addr"`
		Flow_server_ip_subnet_maskSize           int      `json:"flow.server.ip.subnet.mask_size,omitempty"`
		Flow_Server_L4_Port_ID                   int      `json:"flow.server.l4.port.id,omitempty"`
		Flow_Server_L4_Port_Name                 string   `json:"flow.server.l4.port.name,omitempty"`
		Flow_Server_Mac_Addr                     string   `json:"flow.server.mac.addr,omitempty"`
		Flow_Src_As_Asn                          int      `json:"flow.src.as.asn"`
		Flow_Src_As_Label                        string   `json:"flow.src.as.label"`
		Flow_Src_As_Org                          string   `json:"flow.src.as.org"`
		Flow_Src_Host_Name                       string   `json:"flow.src.host.name"`
		Flow_Src_Ip_Addr                         string   `json:"flow.src.ip.addr"`
		Flow_src_ip_subnet_maskSize              int      `json:"flow.src.ip.subnet.mask_size,omitempty"`
		Flow_Src_L4_Port_ID                      int      `json:"flow.src.l4.port.id,omitempty"`
		Flow_Src_L4_Port_Name                    string   `json:"flow.src.l4.port.name,omitempty"`
		Flow_Src_Mac_Addr                        string   `json:"flow.src.mac.addr,omitempty"`
		Gre_Version_Ver                          int      `json:"gre.version.ver,omitempty"`
		Icmp_Code_Name                           string   `json:"icmp.code.name,omitempty"`
		Icmp_ID                                  int      `json:"icmp.id,omitempty"`
		Icmp_seqNum                              int      `json:"icmp.seq_num,omitempty"`
		Icmp_Type_Name                           string   `json:"icmp.type.name,omitempty"`
		Ip_Dscp_Name                             string   `json:"ip.dscp.name"`
		Ip_Ecn_Name                              string   `json:"ip.ecn.name"`
		Ip_Frag_Flags_Tags                       []string `json:"ip.frag.flags.tags,omitempty"`
		Ip_Packet_Size                           int      `json:"ip.packet.size"`
		Ip_Ttl                                   int      `json:"ip.ttl"`
		Ip_v6_flowLabel                          int      `json:"ip.v6.flow_label,omitempty"`
		Ip_Version_Name                          string   `json:"ip.version.name"`
		Ip_Version_Ver                           int      `json:"ip.version.ver"`
		L2_Frame_Size                            int      `json:"l2.frame.size"`
		L4_Proto_Name                            string   `json:"l4.proto.name"`
		L4_Session_Established                   string   `json:"l4.session.established,omitempty"`
		Sflow_Pen_ID                             int      `json:"sflow.pen.id"`
		Sflow_Pen_Name                           string   `json:"sflow.pen.name"`
		Sflow_sample_headerProto_name            string   `json:"sflow.sample.header_proto.name"`
		Sflow_sample_seqNum                      int      `json:"sflow.sample.seq_num"`
		Sflow_Sample_Size                        int      `json:"sflow.sample.size"`
		Sflow_sample_stripSize                   int      `json:"sflow.sample.strip_size"`
		Sflow_sampleType_name                    string   `json:"sflow.sample_type.name"`
		Sflow_sourceID                           int      `json:"sflow.source_id"`
		Sflow_sourceIDType_name                  string   `json:"sflow.source_id_type.name"`
		Sflow_subAgentID                         int      `json:"sflow.sub_agent_id"`
		Tcp_ackNum                               int      `json:"tcp.ack_num,omitempty"`
		Tcp_Flags_Bits                           int      `json:"tcp.flags.bits,omitempty"`
		Tcp_Flags_Tags                           []string `json:"tcp.flags.tags,omitempty"`
		Tcp_Header_Size                          int      `json:"tcp.header.size,omitempty"`
		Tcp_Options_Payload                      string   `json:"tcp.options.payload,omitempty"`
		Tcp_seqNum                               int      `json:"tcp.seq_num,omitempty"`
		Tcp_urgentPointer                        int      `json:"tcp.urgent_pointer,omitempty"`
		Tcp_Window_Size                          int      `json:"tcp.window.size,omitempty"`
		Tunnel_Dst_As_Asn                        int      `json:"tunnel.dst.as.asn,omitempty"`
		Tunnel_Dst_As_Label                      string   `json:"tunnel.dst.as.label,omitempty"`
		Tunnel_Dst_As_Org                        string   `json:"tunnel.dst.as.org,omitempty"`
		Tunnel_Dst_Host_Name                     string   `json:"tunnel.dst.host.name,omitempty"`
		Tunnel_Dst_Ip_Addr                       string   `json:"tunnel.dst.ip.addr,omitempty"`
		Tunnel_Dst_Mac_Addr                      string   `json:"tunnel.dst.mac.addr,omitempty"`
		Tunnel_Ip_Dscp_Name                      string   `json:"tunnel.ip.dscp.name,omitempty"`
		Tunnel_Ip_Ecn_Name                       string   `json:"tunnel.ip.ecn.name,omitempty"`
		Tunnel_Ip_Frag_Flags_Tags                []string `json:"tunnel.ip.frag.flags.tags,omitempty"`
		Tunnel_Ip_Packet_Size                    int      `json:"tunnel.ip.packet.size,omitempty"`
		Tunnel_Ip_Ttl                            int      `json:"tunnel.ip.ttl,omitempty"`
		Tunnel_ip_v6_flowLabel                   int      `json:"tunnel.ip.v6.flow_label,omitempty"`
		Tunnel_Ip_Version_Name                   string   `json:"tunnel.ip.version.name,omitempty"`
		Tunnel_Ip_Version_Ver                    int      `json:"tunnel.ip.version.ver,omitempty"`
		Tunnel_L4_Proto_Name                     string   `json:"tunnel.l4.proto.name,omitempty"`
		Tunnel_Src_As_Asn                        int      `json:"tunnel.src.as.asn,omitempty"`
		Tunnel_Src_As_Label                      string   `json:"tunnel.src.as.label,omitempty"`
		Tunnel_Src_As_Org                        string   `json:"tunnel.src.as.org,omitempty"`
		Tunnel_Src_Host_Name                     string   `json:"tunnel.src.host.name,omitempty"`
		Tunnel_Src_Ip_Addr                       string   `json:"tunnel.src.ip.addr,omitempty"`
		Tunnel_Src_Mac_Addr                      string   `json:"tunnel.src.mac.addr,omitempty"`
		Tunnel_Type_Name                         string   `json:"tunnel.type.name,omitempty"`
		Tunnel_vlan_cTag_dei_state               string   `json:"tunnel.vlan.c_tag.dei.state,omitempty"`
		Tunnel_vlan_cTag_id                      int      `json:"tunnel.vlan.c_tag.id,omitempty"`
		Tunnel_vlan_cTag_pcp_name                string   `json:"tunnel.vlan.c_tag.pcp.name,omitempty"`
		Udp_Size                                 int      `json:"udp.size,omitempty"`
		Vlan_cTag_dei_state                      string   `json:"vlan.c_tag.dei.state,omitempty"`
		Vlan_cTag_id                             int      `json:"vlan.c_tag.id,omitempty"`
		Vlan_cTag_pcp_name                       string   `json:"vlan.c_tag.pcp.name,omitempty"`
		Vlan_Tag_ID                              []int    `json:"vlan.tag.id"`
	} `json:"_source"`
	Fields struct {
		Timestamp              []time.Time `json:"@timestamp"`
		Flow_Collect_Timestamp []time.Time `json:"flow.collect.timestamp"`
	} `json:"fields"`
}