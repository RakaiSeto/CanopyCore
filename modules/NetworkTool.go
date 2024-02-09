package modules

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/valyala/fasthttp"
	_ "github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func HTTPConvertHeaderToMap(httpHeader http.Header) map[string]interface{} {
	var respMap = make(map[string]interface{})

	for key, val := range httpHeader {
		respMap[key] = val
	}

	return respMap
}

func HTTPGETString(traceCode string, mainUrl string, headerRequest map[string]interface{}, urlParameter map[string]string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_GET_String", fmt.Sprintf("HTTP GET - body: %v, header: %v", urlParameter, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("GET", mainUrl, nil)

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_GET_String",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		// Buat encoded url parameter
		query := url.Values{}

		if urlParameter != nil {
			for key, val := range urlParameter {
				query.Add(key, val)
			}
		}
		req.URL.RawQuery = query.Encode()

		// Buat HTTP Webclient
		client := &http.Client{
			Timeout: 30 * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)
		req.Close = true // Close once done

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_GET_String",
				"Failed to execute HTTP Request. Error occured.", true, err)
		} else {
			endDateTime = time.Now()

			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			// fill bodyResponse, headerResponse, dll
			bodyResponse = string(body)
			headerResponse = HTTPConvertHeaderToMap(resp.Header)
			httpStatus = resp.Status
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = req.URL.RawQuery
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_GET_String",
		fmt.Sprintf("mapResponse: %v", mapResponse), true, err)

	return mapResponse
}

//noinspection GoUnusedExportedFunction
func HTTPPOSTString(traceCode string, url string, headerRequest map[string]interface{}, bodyRequest string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_POST_String", fmt.Sprintf("HTTP POST - body: %s, header: %v", bodyRequest, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	var bodyByte = []byte(bodyRequest)

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_POST_String",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		client := &http.Client{
			Timeout: 30 * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)
		req.Close = true // Close once done

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_POST_String",
				"Failed to execute HTTP Request. Error occured.", true, err)
		} else {
			endDateTime = time.Now()

			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			// fill bodyResponse, headerResponse, dll
			bodyResponse = string(body)
			headerResponse = HTTPConvertHeaderToMap(resp.Header)
			httpStatus = resp.Status
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = bodyRequest
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTP_POST_String",
		fmt.Sprintf("mapResponse: %v", mapResponse), true, err)

	return mapResponse
}

//noinspection GoUnusedExportedFunction
func HTTPSPOSTString(traceCode string, url string, headerRequest map[string]interface{}, bodyRequest string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String", fmt.Sprintf("HTTP POST - body: %s, header: %v", bodyRequest, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	var bodyByte = []byte(bodyRequest)

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxConnsPerHost:       1000,
	}

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		client := &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)
		req.Close = true // Close once done

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String",
				"Failed to execute HTTP Request. Error occured.", true, errR)
		} else {
			endDateTime = time.Now()

			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			// fill bodyResponse, headerResponse, dll
			bodyResponse = string(body)
			headerResponse = HTTPConvertHeaderToMap(resp.Header)
			httpStatus = resp.Status
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = bodyRequest
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String",
		fmt.Sprintf("mapResponse: %v", mapResponse), true, err)

	return mapResponse
}

func HTTPSPatchString(traceCode string, url string, headerRequest map[string]interface{}, bodyRequest string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_PATCH_String", fmt.Sprintf("HTTP POST - body: %s, header: %v", bodyRequest, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	var bodyByte = []byte(bodyRequest)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyByte))

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_PATCH_String",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		client := &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_PATCH_String",
				"Failed to execute HTTP Request. Error occured.", true, err)
		} else {
			endDateTime = time.Now()

			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			// fill bodyResponse, headerResponse, dll
			bodyResponse = string(body)
			headerResponse = HTTPConvertHeaderToMap(resp.Header)
			httpStatus = resp.Status
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = bodyRequest
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_PATCH_String",
		fmt.Sprintf("mapResponse: %v", mapResponse), true, err)

	return mapResponse
}

//noinspection GoUnusedExportedFunction
func HTTPSGETString(traceCode string, mainUrl string, headerRequest map[string]interface{}, urlParameter map[string]string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String", fmt.Sprintf("HTTPS GET - body: %v, header: %v", urlParameter, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("GET", mainUrl, nil)

	if req != nil {
		req.Close = true // Close once done
	}

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		// Buat encoded url parameter
		query := url.Values{}

		if urlParameter != nil {
			for key, val := range urlParameter {
				query.Add(key, val)
			}
		}
		req.URL.RawQuery = query.Encode()

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String",
			fmt.Sprintf("Request URL: %v", req.URL.String()), false, nil)

		// Buat HTTP Webclient
		client := &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String",
				"Failed to execute HTTP Request. Error occured.", true, err)
		} else {
			endDateTime = time.Now()

			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			// fill bodyResponse, headerResponse, dll
			bodyResponse = string(body)
			headerResponse = HTTPConvertHeaderToMap(resp.Header)
			httpStatus = resp.Status
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = req.URL.RawQuery
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String",
		fmt.Sprintf("mapResponse: %v", mapResponse), true, err)

	return mapResponse
}

//noinspection GoUnusedExportedFunction
func HTTPSPOSTForm(traceCode string, httpUrl string, headerRequest map[string]interface{}, mapFormContent map[string]string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_FORM", fmt.Sprintf("HTTP POST - body: %s, header: %v", mapFormContent, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	// make form data
	form := url.Values{}

	for key, val := range mapFormContent {
		form.Add(key, val)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(form.Encode()))

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_FORM",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Close = true

		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		client := &http.Client{
			Transport: tr,
			Timeout:   10 * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_FORM",
				"Failed to execute HTTP Request. Error occured.", true, errR)
		} else {
			endDateTime = time.Now()

			body, _ := ioutil.ReadAll(resp.Body)

			// fill bodyResponse, headerResponse, dll
			bodyResponse = string(body)
			headerResponse = HTTPConvertHeaderToMap(resp.Header)
			httpStatus = resp.Status
		}

		if resp != nil {
			resp.Body.Close()
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = fmt.Sprintf("%+v", mapFormContent)
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_FORM",
		fmt.Sprintf("mapResponse: %v", mapResponse), true, err)

	return mapResponse
}

//noinspection GoUnusedExportedFunction
func HTTPSGETDownloadFile(traceCode string, fileUrl string, headerRequest map[string]interface{}, targetDirectory string, targetFileName string, downloadingTimeOut int) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE", fmt.Sprintf("HTTP GET - header: %v", headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse map[string]interface{}
	var bodyResponse []byte
	var fileName string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Compose request with body
	startDateTime = time.Now()
	req, err := http.NewRequest("GET", fileUrl, nil)

	if err != nil {
		hitStatus = "240"

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
			"Failed to compose HTTP Request. Error occured.", true, err)
	} else {
		// Compose header
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
				DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
					"Header - "+key+" - "+val.(string), false, nil)
			}
		}

		// Buat HTTP Webclient
		//client := &http.Webclient{Transport: tr}

		// Get the response
		client := &http.Client{
			Transport: tr,
			Timeout:   time.Duration(downloadingTimeOut) * time.Second,
		}

		// Get the response
		resp, errR := client.Do(req)

		if errR != nil {
			hitStatus = "240"

			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
				"Failed to execute HTTP Request. Error occured.", true, errR)
		} else {
			endDateTime = time.Now()

			defer resp.Body.Close()

			// Create targetFile
			fullPathTargetFile := targetDirectory + targetFileName
			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
				"fullPathTargetFile: "+fullPathTargetFile, false, nil)

			file, errFile := os.Create(fullPathTargetFile)

			if errFile != nil {
				hitStatus = "240"
				DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
					"Failed to save the file. Error occured.", true, errFile)
			} else {
				defer file.Close()

				// Copy file to target
				_, errCopy := io.Copy(file, resp.Body)

				if errCopy != nil {
					hitStatus = "240"
					DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
						"Failed to save the file. Error occured.", true, errCopy)
				} else {
					//body, _ := ioutil.ReadAll(resp.Body)

					// fill bodyResponse, headerResponse, dll
					//bodyResponse = body
					headerResponse = HTTPConvertHeaderToMap(resp.Header)
					httpStatus = resp.Status
				}
			}
		}
	}

	timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

	// Put into mapResponse
	mapResponse["bodyRequest"] = fileUrl + "?" + req.URL.RawQuery
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["fileName"] = fileName
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_DOWNLOAD_FILE",
		fmt.Sprintf("mapResponse: %v", mapResponse), false, err)

	return mapResponse
}

func HTTPSPOSTStringFast(traceCode string, url string, headerRequest map[string]interface{}, bodyRequest string) map[string]interface{} {
	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String_FAST",
		fmt.Sprintf("HTTP POST - body: %s, header: %v", bodyRequest, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse = make(map[string]interface{})
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.SetBody([]byte(bodyRequest))
	req.Header.SetMethod("POST") // Method

	// Set Headers
	if headerRequest != nil {
		for key, val := range headerRequest {
			req.Header.Set(key, val.(string))
		}
	}

	// Prepare response
	res := fasthttp.AcquireResponse()

	// Do hit
	if err := fasthttp.Do(req, res); err != nil {
		// Release request
		fasthttp.ReleaseRequest(req)

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String_FAST",
			"Failed hit remote URL. Error occur.", true, err)

		hitStatus = "901"
	} else {
		// Proses
		endDateTime = time.Now()
		timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

		// Release request
		fasthttp.ReleaseRequest(req)

		bodyResponse = string(res.Body())
		origheaderResponse := res.Header
		httpStatus = strconv.Itoa(res.StatusCode())

		origheaderResponse.VisitAll(func(key, value []byte) {
			headerResponse[string(key)] = string(value)
		})

		// Release response
		fasthttp.ReleaseResponse(res)

		hitStatus = "000"
	}

	// Put into mapResponse
	mapResponse["bodyRequest"] = bodyRequest
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus
	mapResponse["url"] = url

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_POST_String_FAST",
		fmt.Sprintf("mapResponse: %v", mapResponse), false, nil)

	return mapResponse
}

func HTTPSGETStringFast(traceCode string, theUrl string, headerRequest map[string]interface{}, mapRequest map[string]string) map[string]interface{} {
	//DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String_FAST",
	//	fmt.Sprintf("HTTP POST - body: %+v, header: %v", mapRequest, headerRequest), false, nil)
	var mapResponse = make(map[string]interface{})
	var headerResponse = make(map[string]interface{})
	var bodyResponse string
	var startDateTime = time.Now()
	var endDateTime = time.Now()
	var timeDiff int64 // millisecond
	var httpStatus string
	var hitStatus = "000"

	// Convert mapRequest to GET string query
	base, err := url.Parse(theUrl)

	if err != nil {
		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String_FAST",
			"Failed to create complete GET URL. Error occur.", true, err)
	} else {
		// Compose query string
		params := url.Values{}

		if mapRequest != nil {
			for keyPar, valPar := range mapRequest {
				params.Add(keyPar, valPar)
			}
		}
		base.RawQuery = params.Encode()
		theUrl = base.String()

		DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String_FAST",
			"HTTP GET complete url: "+theUrl, false, nil)

		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.SetRequestURI(theUrl)

		// Set Headers
		if headerRequest != nil {
			for key, val := range headerRequest {
				req.Header.Set(key, val.(string))
			}
		}

		// Prepare response
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(res)

		// Do hit
		if errX := fasthttp.Do(req, res); errX != nil {
			DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String",
				"Failed hit remote URL. Error occur.", true, errX)

			hitStatus = "901"
		} else {
			// Proses
			endDateTime = time.Now()
			timeDiff = endDateTime.Sub(startDateTime).Milliseconds()

			// Release request
			fasthttp.ReleaseRequest(req)

			bodyResponse = string(res.Body())
			//origheaderResponse := res.Header
			httpStatus = strconv.Itoa(res.StatusCode())

			//origheaderResponse.VisitAll(func(key, value []byte) {
			//	headerResponse[string(key)] = string(value)
			//})

			hitStatus = "000"
		}
	}

	// Put into mapResponse
	mapResponse["bodyRequest"] = theUrl
	mapResponse["headerRequest"] = headerRequest
	mapResponse["bodyResponse"] = bodyResponse
	mapResponse["headerResponse"] = headerResponse
	mapResponse["httpStatus"] = httpStatus
	mapResponse["startDateTime"] = startDateTime
	mapResponse["endDateTime"] = endDateTime
	mapResponse["timeDifference"] = timeDiff
	mapResponse["status"] = hitStatus
	mapResponse["url"] = theUrl

	DoLog("DEBUG", traceCode, "NetworkTool", "HTTPS_GET_String",
		fmt.Sprintf("mapResponse: %v", mapResponse), false, nil)

	return mapResponse
}
