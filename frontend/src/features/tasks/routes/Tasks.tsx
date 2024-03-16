import { ContentLayout } from '@/components/Layout';
import { TaskList } from '../components/TasksList';
import { CreateDialog } from '../components/CreateDialog';
import storage from '@/util/storage';

export const Tasks = () => {
  //TODO: 認証や全体レイアウトを追加

  const userId = storage.getUserId();

  return (
    <ContentLayout title="タスク一覧">
      <div className="mt-4">
        <TaskList userId={userId} />
      </div>
      <div className="mt-4">
        <CreateDialog userId={userId} />
      </div>
    </ContentLayout>
  );
};
