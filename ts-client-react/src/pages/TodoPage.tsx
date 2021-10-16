import React, { useState } from "react";
// import { Heading, Grid, Box, Flex, Text, CloseButton, Button, Input } from "@chakra-ui/react";
import { Input, Button, Heading } from "../material-tailwind";
import { useTodos } from "../hooks/todo";
import { TodoItem } from "../rpc/todo";

function Item({ todo }: { todo: TodoItem }) {
  const { deleteTodo } = useTodos();
  const { task, id } = todo;

  return (
    <li className="bg-white text-center p-2 m-1 w-2/6 border border-color-gray-500">
      <div className="flex justify-between">
        <h5 className="font-light">{task}</h5>
        <button
          onClick={() => deleteTodo(id)}
          className="text-red-500 pl-2 pr-2 font-medium rounded-xl hover:bg-red-600 hover:text-white"
        >
          X
        </button>
      </div>
    </li>
  );
}

function List({ todos }: { todos: TodoItem[] }) {
  return (
    <ul className="grid justify-items-center">
      {todos.map((todo) => (
        <Item key={todo.id} todo={todo} />
      ))}
    </ul>
  );
}

// function Item({ todo }: { todo: TodoItem }) {
//   const { deleteTodo } = useTodos();
//   const { task, id } = todo;

//   return (
//     <Box w={2 / 6} border="1px" borderColor="gray.200" margin="1" padding="2">
//       <Flex justifyContent="space-between" alignItems="baseline">
//         <Text>{task}</Text>
//         <CloseButton size="sm" color="red.500" onClick={() => deleteTodo(id)} />
//       </Flex>
//     </Box>
//   );
// }

// function List({ todos }: { todos: TodoItem[] }) {
//   return (
//     <Grid justifyItems="center">
//       {todos.map((todo) => (
//         <Item key={todo.id} todo={todo} />
//       ))}
//     </Grid>
//   );
// }

export function TodoPage() {
  const { addTodo, todos } = useTodos();
  const [addText, setAddText] = useState("");

  function handleAdd() {
    addTodo(addText);
    setAddText("");
  }

  return (
    <div id="app" className="font-sans">
      <section>
        <Heading
          as="h3"
          color="gray"
          size="2xl"
          className="text-gray-800 text-center max-w-2xl m-auto font-light mt-5 mb-5 p-5"
        >
          TODO gRPC Client
        </Heading>
        {/* <h3 className="text-gray-800 text-center max-w-2xl text-2xl font-light m-auto mt-5 mb-5 p-5">
        </h3> */}
        <form
          className="flex justify-center"
          onSubmit={(e) => {
            e.preventDefault();
            handleAdd();
          }}
        >
          <div className="flex space-x-5" role="add">
            <Input
              placeholder="Walk my dog"
              // className="pl-2 p-2 border-b border-gray-300 rounded-md w-64"
              // className="pl-2 p-2 border-b border-gray-300 rounded-md w-64"
              type="text"
              value={addText}
              onChange={(e) => setAddText(e.target.value)}
            />
            <Button
              className="text-base whitespace-nowrap"
              // className="hover:bg-blue-800 bg-blue-600 rounded-md text-white px-5 text-base py-1"
              size="sm"
              type="submit"
              name="action"
              onClick={(e) => {
                e.preventDefault();
                handleAdd();
              }}
            >
              Add Todo
            </Button>
          </div>
        </form>
      </section>
      <div className="mt-6">
        <List todos={todos} />
      </div>
    </div>
  );

  // return (
  //   <div>
  //     <section>
  //       <Heading
  //         as="h3"
  //         size="xl"
  //         textColor="gray.800"
  //         textAlign="center"
  //         fontWeight="light"
  //         marginTop="5"
  //         marginBottom="5"
  //         padding="5"
  //       >
  //         TODO gRPC Client
  //       </Heading>
  //       <form
  //         onSubmit={(e) => {
  //           e.preventDefault();
  //           handleAdd();
  //         }}
  //       >
  //         <Flex justifyContent="center">
  //           <Flex>
  //             <Input
  //               padding="2"
  //               width="80"
  //               placeholder="Walk my dog"
  //               value={addText}
  //               type="text"
  //               onChange={(e) => setAddText(e.target.value)}
  //             />
  //             <Button
  //               marginLeft="5"
  //               size="md"
  //               colorScheme="blue"
  //               variant="solid"
  //               onClick={(e) => {
  //                 e.preventDefault();
  //                 handleAdd();
  //               }}
  //             >
  //               Add Todo
  //             </Button>
  //           </Flex>
  //         </Flex>
  //       </form>
  //     </section>
  //     <Box marginTop="6">
  //       <List todos={todos} />
  //     </Box>
  //   </div>
  // );
}
