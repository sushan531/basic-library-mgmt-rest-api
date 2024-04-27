package utils

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type PageType string

const (
	PageTypeNextPage     PageType = "NextPage"
	PageTypePreviousPage PageType = "PreviousPage"
)

func DecodePaginationCursor(cursor string) (PageType, int32, int32, error) {
	decodedCursor, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		return "NextPage", 0, 0, err
	}
	// Split the decoded cursor into parts using "-" as the separator
	parts := strings.Split(string(decodedCursor), "||")
	if len(parts) != 3 {
		return "NextPage", 0, 0, fmt.Errorf("invalid pagination cursor format")
	}
	// Parse command
	command := parts[0]

	searchStringStart := parts[1]
	searchStringEnd := parts[2]
	searchIdStart, _ := strconv.ParseInt(searchStringStart, 10, 32)
	searchIdEnd, _ := strconv.ParseInt(searchStringEnd, 10, 32)
	return PageType(command), int32(searchIdStart), int32(searchIdEnd), nil
}

func EncodePaginationCursor(command PageType, startId int32, endId int32) string {
	cursor := fmt.Sprintf("%s||%d||%d", command, startId, endId)
	encodedCursor := base64.URLEncoding.EncodeToString([]byte(cursor))
	return encodedCursor
}
