import React, { useState } from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { MenuItem } from '@mui/material';
import { useCreateTask } from '../api/createTask';

//TODO: 一箇所で定義する
const status = [
  {
    value: 'NOT_STARTED',
    label: 'NOT_STARTED',
  },
  {
    value: 'IN_PROGRESS',
    label: 'IN_PROGRESS',
  },
  {
    value: 'DONE',
    label: 'DONE',
  },
];

type Props = {
  userId: string;
};

export const CreateDialog = (props: Props) => {
  const initialNewTask = {
    userId: props.userId,
    title: '',
    description: '',
    status: 'NOT_STARTED',
  };
  const [open, setOpen] = useState(false);
  const [task, setTask] = useState(initialNewTask);

  const taskMutation = useCreateTask();

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
    setTask(initialNewTask);
  };

  const handleSubmit = () => {
    //TODO: 処理失敗時の挙動を追加する
    taskMutation.mutate({ userId: props.userId, data: task });
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setTask((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  return (
    <React.Fragment>
      <Button variant="outlined" onClick={handleClickOpen}>
        タスク新規作成
      </Button>
      <Dialog
        open={open}
        onClose={handleClose}
        PaperProps={{
          component: 'form',
          onSubmit: (event: React.FormEvent<HTMLFormElement>) => {
            event.preventDefault();
            handleClose();
          },
        }}
      >
        <DialogTitle>タスクの編集</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            required
            margin="dense"
            name="title"
            label="タイトル"
            type="text"
            fullWidth
            variant="standard"
            value={task.title}
            onChange={handleInputChange}
          />
          <TextField
            autoFocus
            required
            margin="dense"
            name="description"
            label="説明"
            type="text"
            fullWidth
            variant="standard"
            value={task.description}
            onChange={handleInputChange}
          />
          <TextField
            autoFocus
            required
            margin="dense"
            name="status"
            type="text"
            fullWidth
            variant="standard"
            value={task.status}
            select
            label="ステータス"
            defaultValue={task.status}
            onChange={handleInputChange}
          >
            {status.map((option) => (
              <MenuItem key={option.value} value={option.value}>
                {option.label}
              </MenuItem>
            ))}
          </TextField>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>キャンセル</Button>
          <Button onClick={handleSubmit} type="submit">
            作成
          </Button>
        </DialogActions>
      </Dialog>
    </React.Fragment>
  );
};
