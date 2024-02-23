export type Task = {
  taskId: string;
  userId: string;
  title: string;
  description: string;
  status: string;
};

export type NewTask = Omit<Task, 'taskId'>;
