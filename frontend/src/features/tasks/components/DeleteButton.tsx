import { Button } from '@/components/Elements';
import { useDeleteTask } from '../api/deleteTask';

type Props = {
  userId: string;
  taskId: string;
};

export const DeleteButton = (props: Props) => {
  const deleteTask = useDeleteTask();

  const handleOnClick = () => {
    deleteTask.mutate({ userId: props.userId, taskId: props.taskId });
  };

  return (
    <Button size="sm" variant="danger" onClick={handleOnClick}>
      削除
    </Button>
  );
};
