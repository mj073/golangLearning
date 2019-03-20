package ip

type IPHeader struct {
	VersionHLen uint8
	ServiceType uint8
	TotalLength uint16
	Identification uint16
	FlagFragmentationBits uint16
	TTL uint8
	Protocol uint8
	HeaderChecksum uint16
	SourceIP	uint32
	DestinationIP 	uint32
	Option [40]byte
}

func main()  {
}