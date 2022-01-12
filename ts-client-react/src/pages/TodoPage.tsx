import React from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { Heading, Box, Text, CloseButton, Grid, Flex, Input, Button } from "@chakra-ui/react";
import { useTodos } from "../hooks/todo";
import { TodoItem } from "../rpc/todo";
import { useAuth } from "../hooks/auth";

type FormFields = {
  text: string;
};

function Item({ todo }: { todo: TodoItem }) {
  const { mutations } = useTodos();
  const [deleteTodo] = mutations.useDeleteTodo();
  const { task, id } = todo;

  return (
    <Box w={2 / 6} border="1px" borderColor="gray.200" margin="1" padding="2">
      <Flex justifyContent="space-between" alignItems="baseline">
        <Text>{task}</Text>
        <CloseButton size="sm" color="red.500" onClick={() => deleteTodo({ id })} />
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
  const { mutations, todos } = useTodos();
  const [addTodo] = mutations.useAddTodo();
  const { register, handleSubmit, setValue } = useForm<FormFields>();
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    navigate("/auth/login");
    return null;
  }

  function onSubmit(data: FormFields) {
    addTodo({ task: data.text });
    setValue("text", "");
  }

  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <Box w="100%" bgColor="white" p="5">
        <Heading as="h3" size="xl" textColor="gray.800" textAlign="center" fontWeight="light" padding="5">
          TODO gRPC Client
        </Heading>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Flex justifyContent="center">
            <Flex>
              <Input padding="2" width="80" placeholder="Walk my dog" type="text" {...register("text")} />
              <Button marginLeft="5" size="md" colorScheme="blue" variant="solid" onClick={handleSubmit(onSubmit)}>
                Add Todo
              </Button>
            </Flex>
          </Flex>
        </form>
      </Box>
      <Box p="5" bgColor="white">
        <List todos={todos} />
      </Box>
    </Box>
  );
}
