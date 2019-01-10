const http = require('http');
const fs = require('fs');
const hostname = '127.0.0.1';
const port = 3000;

// fs.readFile('index.html', (err, html) => {
//      if (err) {
//          throw err;
//      }

//     const server = http.createServer((requestAnimationFrame, res) => {
//         res.statusCode = 200;
//         res.setHeader('Content-type', 'text/html');
//         res.write(html);
//         res.end();
//     });

//     server.listen(port, hostname, () => {
//         console.log('Server started on port ' + port);
//     })

// });


const server = http.createServer((requestAnimationFrame, res) => {
    res.statusCode = 200;
    res.setHeader('Content-type', 'text/html');
    res.write(html);
    res.end();
});

server.listen(port, hostname, () => {
    console.log('Server started on port ' + port);
})

function greeting(greet) {
    console.log(greet);
}


let req = {
    method: 'GET',
    responseType: 'arraybuffer',
    url: 'my-protobuf-endpoint'
}

// const Vehicle = proto.main.Vehicle;

// return $http(req).then(function(response) {
//     // We need to encapsulate the response on Uint8Array to avoid
//     // getting it converted to string.
//     ctrl.vehicle = Vehicle.decode(new Uint8Array(response.data)).vehicle;
//   });



// Issuing the POST request built above.
// return $http(post).then(function() {
//     console.log('Everything went just fine');
//   });