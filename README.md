**sentry-slack-bridge**

    go build . 
    ./sentry-slack-bridge 
    curl -v -H "Content-Type: application/json" -X POST --data @post.json http://localhost:8080/op

**config.yaml example**

    - name: first  
	    options:  
		    path: "first"  
		    channel: "first"  
		    webhook: "https://hooks.slack.com/services/11111/1111111/bGkghRj36nA80t6Rg94XOzgB"  
      
    - name: second  
	    options:  
		    path: "second"  
		    channel: "second"  
		    webhook: "https://hooks.slack.com/services/11111/1111111/xJkMfTjtSnA80t6Rg94XOzOY"

**Docker**

    docker run -it -p 8080:8080 -v $(pwd)/config.yaml:/app/config.yaml pontostroy/sentry-slack-bridge

