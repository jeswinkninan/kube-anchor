package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

	v1beta1 "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

//HandleFreeze Func will implementing the http handler for handling client calls to /freeze
func HandleFreeze(responseWriter http.ResponseWriter, request *http.Request){

	// Reading the admission request body
	body, err := ioutil.ReadAll(request.Body)
	if err !=nil{
		log.Printf("Unable to read the request")
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
	defer request.Body.Close()

	// Decode the request into AdmissionReview struct
	admissionReview := v1beta1.AdmissionReview{}
	err = json.Unmarshal(body, &admissionReview);
	if err !=nil{
		log.Printf("unmarshaling request failed with %s", err)
	}

	//Response Objects Declarations
	admReviewObj := admissionReview.Request	
	anchorResp := v1beta1.AdmissionResponse{}
	anchorResp.UID = admReviewObj.UID
	anchorResp.Allowed = false
	admissionReview.Response = &anchorResp
	anchorResp.Result = &metav1.Status{
		Status: "Failure",
		Message: "Cluster Freeze Window Enabled via Kube-Anchor ☸ ",
		Code: 403,
	}

	// Encoding the response body with objects referenced
	anchorResponseBody, err := json.Marshal(admissionReview)
	if err != nil {
			log.Printf("Unable to encode the response %s", err)
	}

	//Writing back as Response
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(anchorResponseBody)

}

func main() {

	http.HandleFunc("/freeze", HandleFreeze)
	err := http.ListenAndServeTLS(":443", "./cert/tls.crt", "./cert/tls.key", nil)
	if err != nil {
		log.Fatalf("Failed to start the kube-anchor webhook server :%s", err)
	}
	log.Printf("Started the webhook server")

}