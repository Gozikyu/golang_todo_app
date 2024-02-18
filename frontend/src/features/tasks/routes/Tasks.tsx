import { ContentLayout } from '@/components/Layout';
import { TaskList } from '../components/TasksList';

export const Tasks = () => {
  //TODO: 認証や全体レイアウトを追加

  return (
    <ContentLayout title="タスク一覧">
      <TaskList />
    </ContentLayout>
  );
};
