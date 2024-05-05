package utils

import "github.com/bwmarrin/snowflake"

type IdGenerator interface {
	GenerateId() int64
}

type snowflakeIdGenerator struct {
	node *snowflake.Node
}

func NewSnowflakeIdGenerator(node *snowflake.Node) IdGenerator {
	return &snowflakeIdGenerator{
		node: node,
	}
}

func (gen snowflakeIdGenerator) GenerateId() int64 {
	return gen.node.Generate().Int64()
}
