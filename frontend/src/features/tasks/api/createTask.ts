import { axios } from '@/lib/axios';
import { useMutation } from 'react-query';
import { NewTask } from '../types';
import { queryClient } from '@/lib/react-query';

export type CreateTaskVariables = {
  userId: string;
  data: NewTask;
};

const createTask = ({ userId, data }: CreateTaskVariables) => {
  return axios.post(`/${userId}/tasks`, data);
};

export const useCreateTask = () => {
  return useMutation({
    mutationFn: createTask,
    onSuccess: () => {
      queryClient.invalidateQueries('tasks');
    },
    onError: (error) => {
      console.log(`mutation error ${error}`);
    },
  });
};
