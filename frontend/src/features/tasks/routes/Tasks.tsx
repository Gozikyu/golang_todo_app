import { ContentLayout } from '@/components/Layout';
import { TaskList } from '../components/TasksList';
import { CreateDialog } from '../components/CreateDialog';

export const Tasks = () => {
  //TODO: 認証や全体レイアウトを追加

  return (
    <ContentLayout title="タスク一覧">
      <div className="mt-4">
        <TaskList />
      </div>
      <div className="mt-4">
        {/**TODO: userIdを動的に設定 */}
        <CreateDialog userId="1" />
      </div>
    </ContentLayout>
  );
};
