import { axios } from '@/lib/axios';
import { useMutation } from 'react-query';
import { Task } from '../types';
import { queryClient } from '@/lib/react-query';

export type UpdateTaskVariables = {
  userId: string;
  data: Task;
};

const updateTask = ({ userId, data }: UpdateTaskVariables) => {
  return axios.put(`/${userId}/tasks/${data.taskId}`, data);
};

export const useUpdateTask = () => {
  return useMutation({
    mutationFn: updateTask,
    onSuccess: () => {
      queryClient.invalidateQueries('tasks');
    },
    onError: (error) => {
      console.log(`mutation error ${error}`);
    },
  });
};
