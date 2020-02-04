const http = require('http')

http.createServer(function(request, response) {
	// console.log("URL: " + request.url);
	// console.log("Тип " + request.method);
	// console.log("User-Agent" + request.headers["user-agent"]);
	console.log("Все заголовки");
	console.log(request.headers);
	response.setHeader("Access-Control-Allow-Origin", "*")
	response.end();
}).listen(8080)