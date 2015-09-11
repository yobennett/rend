/**
 * Parses the text-based protocol and returns a command line
 * struct representing the work to be done
 */
package textprot

import "bufio"
import "fmt"
import "io"
import "strconv"
import "strings"

import "../common"

type TextParser struct { }

func (p TextParser) ParseRequest(remoteReader *bufio.Reader) (interface{}, common.RequestType, error) {
    
    data, err := remoteReader.ReadString('\n')
    
    if err != nil {
        if err == io.EOF {
            fmt.Println("End of file: Connection closed?")
        } else {
            fmt.Println(err.Error())
        }
        return nil, common.REQUEST_GET, err
    }
    
    clParts := strings.Split(strings.TrimSpace(data), " ")
    
    switch clParts[0] {
        case "set":
            length, err := strconv.Atoi(strings.TrimSpace(clParts[4]))
            
            if err != nil {
                fmt.Println(err.Error())
                return nil, common.REQUEST_SET, common.BAD_LENGTH
            }
            
            flags, err := strconv.Atoi(strings.TrimSpace(clParts[2]))
            
            if err != nil {
                fmt.Println(err.Error())
                return nil, common.REQUEST_SET, common.BAD_FLAGS
            }
            
            return common.SetRequest {
                Key:     clParts[1],
                Flags:   flags,
                Exptime: clParts[3],
                Length:  length,
            }, common.REQUEST_SET, nil
            
        case "get":
            return common.GetRequest {
                Keys: clParts[1:],
            }, common.REQUEST_GET, nil
            
        case "delete":
            return common.DeleteRequest {
                Key: clParts[1],
            }, common.REQUEST_DELETE, nil
            
        // TODO: Error handling for invalid cmd line
        case "touch":
            return common.TouchRequest {
                Key:     clParts[1],
                Exptime: clParts[2],
            }, common.REQUEST_TOUCH, nil
    }
    
    return nil, common.REQUEST_GET, nil
}