import { Spinner, Table } from '@/components/Elements';
import { Task } from '../types';
import { useTasks } from '../api/getTasks';

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

//TODO: 仮実装 reactQueryを使うようにする
// const queryTasks = (): Task[] => {
//   return dummyTasks;
// };

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

  console.log(tasksQuery.data);

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
