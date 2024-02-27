import { Button } from '@/components/Elements';
import { useDeleteTask } from '../api/deleteTask';

type Props = {
  taskId: string;
};

export const DeleteButton = (props: Props) => {
  const deleteTask = useDeleteTask();

  const handleOnClick = () => {
    deleteTask.mutate({ userId: '1', taskId: props.taskId }); //TODO: userIdを動的に設定
  };

  return (
    <Button size="sm" variant="danger" onClick={handleOnClick}>
      削除
    </Button>
  );
};
