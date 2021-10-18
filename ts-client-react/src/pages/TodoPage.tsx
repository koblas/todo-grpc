import React, { useState } from "react";
import { useHistory } from "react-router-dom";
import { Heading, Box, Text, CloseButton, Grid, Flex, Input, Button } from "@chakra-ui/react";
import { useTodos } from "../hooks/todo";
import { TodoItem } from "../rpc/todo";
import { useAuth } from "../hooks/auth";
import { SidebarWithHeader } from "../components/AppContainer";

function Item({ todo }: { todo: TodoItem }) {
  const { deleteTodo } = useTodos();
  const { task, id } = todo;

  return (
    <Box w={2 / 6} border="1px" borderColor="gray.200" margin="1" padding="2">
      <Flex justifyContent="space-between" alignItems="baseline">
        <Text>{task}</Text>
        <CloseButton size="sm" color="red.500" onClick={() => deleteTodo(id)} />
      </Flex>
    </Box>
  );
}

function List({ todos }: { todos: TodoItem[] }) {
  return (
    <Grid justifyItems="center">
      {todos.map((todo) => (
        <Item key={todo.id} todo={todo} />
      ))}
    </Grid>
  );
}

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
    <SidebarWithHeader>
      <section>
        <Heading
          as="h3"
          size="xl"
          textColor="gray.800"
          textAlign="center"
          fontWeight="light"
          marginTop="5"
          marginBottom="5"
          padding="5"
        >
          TODO gRPC Client
        </Heading>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            handleAdd();
          }}
        >
          <Flex justifyContent="center">
            <Flex>
              <Input
                padding="2"
                width="80"
                placeholder="Walk my dog"
                value={addText}
                type="text"
                onChange={(e) => setAddText(e.target.value)}
              />
              <Button
                marginLeft="5"
                size="md"
                colorScheme="blue"
                variant="solid"
                onClick={(e) => {
                  e.preventDefault();
                  handleAdd();
                }}
              >
                Add Todo
              </Button>
            </Flex>
          </Flex>
        </form>
      </section>
      <Box marginTop="6">
        <List todos={todos} />
      </Box>
    </SidebarWithHeader>
  );
}
