package core


type BlockIndexWriter struct {
	blockMetas []*BlockMeta
	totalBytes int
}

func NewBlockIndexWriter() *BlockIndexWriter {
	return &BlockIndexWriter{
		blockMetas: make([]*BlockMeta, 0),
		totalBytes: 0,
	}
}

func (biw *BlockIndexWriter) append(lastKV KeyValue, offset uint64, size uint64, bloomFilter []byte) {
	meta := NewBlockMeta(lastKV, offset, size, bloomFilter)
	biw.blockMetas = append(biw.blockMetas, meta)
	biw.totalBytes += meta.GetSerializeSize()
}

func (biw *BlockIndexWriter) serialize() []byte {
	buffer := make([]byte, biw.totalBytes)
	pos := 0
	for _, meta := range biw.blockMetas {
		metaBytes := meta.ToBytes()
		//todo
		copy(buffer[pos:pos+len(metaBytes)], metaBytes)
		pos += len(metaBytes)
	}
	return buffer
}
