import { Spinner, Table } from '@/components/Elements';
import { Task } from '../types';
import { useTasks } from '../api/getTasks';

export const TaskList = () => {
  //TODO: userIdを動的に入れる
  const tasksQuery = useTasks('1');

  if (tasksQuery.isLoading) {
    return (
      <div className="w-full h-48 flex justify-center items-center">
        <Spinner size="lg" />
      </div>
    );
  }

  if (!tasksQuery.data) return null;

  return (
    <Table<Task>
      data={tasksQuery.data}
      columns={[
        {
          title: 'タイトル',
          field: 'title',
        },
        {
          title: '説明',
          field: 'description',
        },
        {
          title: 'ステータス',
          field: 'status',
        },
      ]}
    />
  );
};
