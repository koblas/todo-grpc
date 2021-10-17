import React, { useState } from "react";
import DeleteIcon from "@mui/icons-material/Delete";
import { useHistory } from "react-router-dom";
import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { useTodos } from "../hooks/todo";
import { TodoItem } from "../rpc/todo";
import { useAuth } from "../hooks/auth";
import { AppContainer } from "../components/AppContainer";

function Item({ todo }: { todo: TodoItem }) {
  const { deleteTodo } = useTodos();
  const { task, id } = todo;

  return (
    <Box sx={{ border: 1, borderColor: "grey.300", textAlign: "center" }} width={2 / 6} p="0.5rem" m="0.25rem">
      <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
        <Typography variant="body1">{task}</Typography>
        <IconButton aria-label="delete" onClick={() => deleteTodo(id)}>
          <DeleteIcon fontSize="small" />
        </IconButton>
      </Box>
    </Box>
  );
}

function List({ todos }: { todos: TodoItem[] }) {
  return (
    // <Box component="ul" sx={{ display: "grid" }} className="grid justify-items-center">
    <Box sx={{ display: "grid", justifyItems: "center" }}>
      {todos.map((todo) => (
        <Item key={todo.id} todo={todo} />
      ))}
    </Box>
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
  const history = useHistory();
  const { isAuthenticated } = useAuth();
  const [addText, setAddText] = useState("");

  if (!isAuthenticated) {
    history.push("/auth/login");
    return null;
  }

  function handleAdd() {
    addTodo(addText);
    setAddText("");
  }

  return (
    <AppContainer>
      <section>
        <Typography variant="h3" sx={{ textAlign: "center" }} my="1.25rem" p="1.25rem">
          TODO gRPC Client
        </Typography>
        <Box
          component="form"
          sx={{ display: "flex", justifyContent: "center" }}
          onSubmit={(e: React.FormEvent) => {
            e.preventDefault();
            handleAdd();
          }}
        >
          <Box sx={{ display: "flex" }}>
            <TextField
              placeholder="Walk my dog"
              sx={{ width: "32rem", marginRight: "1.25rem" }}
              type="text"
              value={addText}
              onChange={(e) => setAddText(e.target.value)}
            />
            <Button
              variant="contained"
              type="submit"
              onClick={(e) => {
                e.preventDefault();
                handleAdd();
              }}
            >
              Add Todo
            </Button>
          </Box>
        </Box>
      </section>
      <Box pt="1.5rem">
        <List todos={todos} />
      </Box>
    </AppContainer>
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
