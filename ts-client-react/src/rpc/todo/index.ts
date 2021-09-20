export interface TodoItem {
  id: string;
  task: string;
}

export interface TodoService {
  getTodos(): Promise<TodoItem[]>;
  addTodo(task: string): Promise<TodoItem>;
  deleteTodo(id: string): Promise<void>;
}
