const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");
const todoPackageDef = protoLoader.loadSync("todo.proto", {});
const grpcObject = grpc.loadPackageDefinition(todoPackageDef);
const todoPackage = grpcObject.todoPackage;

const argument = process.argv[2];
const todoItem = process.argv[3];

const client = new todoPackage.Todo(
  "localhost:9000",
  grpc.credentials.createInsecure()
);

if (argument === "--read") {
  console.log("Reading TODO Items!");
  client.ReadTodos({}, (err, response) => {
    if (response.items) response.items.forEach((i) => console.log(i.text));
  });
}

if (argument === "--add") {
  console.log("Adding TODO Items!");
  client.CreateTodo({ id: 1, text: todoItem }, (err, response) => {
    console.log("Received from server" + JSON.stringify(response));
  });
}

// const call = client.readTodosStream();
// call.on("data", (item) => {
//   console.log("Received item from server " + JSON.stringify(item));
// });

// call.on("end", (e) => console.log("Server done streaming."));
