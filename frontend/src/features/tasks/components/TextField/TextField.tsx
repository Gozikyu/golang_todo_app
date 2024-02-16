import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import { Task } from '../../types';

const dummyTasks = [
  {
    taskId: 'taskId1',
    userId: 'userId1',
    title: 'title1',
    description: 'description1',
    status: 'done',
  },
  {
    taskId: 'taskId2',
    userId: 'userId2',
    title: 'title2',
    description: 'description2',
    status: 'done',
  },
  {
    taskId: 'taskId3',
    userId: 'userId3',
    title: 'title3',
    description: 'description3',
    status: 'done',
  },
];

//TODO: 仮実装
const queryTasks = (): Task[] => {
  return dummyTasks;
};

type Props = {
  task: Task;
};

export const TextFields = ({ task }: Props) => {
  return (
    <Box
      component="form"
      sx={{
        '& > :not(style)': { m: 1, width: '25ch' },
      }}
      noValidate
      autoComplete="off"
    >
      <TextField
        id="outlined-basic"
        label="title"
        variant="standard"
        value={task.title}
      />
      <TextField
        id="filled-basic"
        label="description"
        variant="standard"
        value={task.description}
      />
      <TextField
        id="standard-basic"
        label="status"
        variant="standard"
        value={task.status}
      />
    </Box>
  );
};

export const TaskRows = () => {
  const tasks = queryTasks();
  return (
    <>
      {tasks.map((task) => (
        <TextFields task={task} />
      ))}
    </>
  );
};
