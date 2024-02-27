import { axios } from '@/lib/axios';
import { useMutation } from 'react-query';
import { queryClient } from '@/lib/react-query';

export type UpdateTaskVariables = {
  userId: string;
  taskId: string;
};

const deleteTask = ({ userId, taskId }: UpdateTaskVariables) => {
  return axios.delete(`/${userId}/tasks/${taskId}`);
};

export const useDeleteTask = () => {
  return useMutation({
    mutationFn: deleteTask,
    onSuccess: () => {
      queryClient.invalidateQueries('tasks');
    },
    onError: (error) => {
      console.log(`mutation error ${error}`);
    },
  });
};
