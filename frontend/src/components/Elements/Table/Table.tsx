import { Task } from '@/features/tasks/types';
import FormDialog from '../FormDialog/FormDialog';
import { DeleteButton } from '@/features/tasks/components/DeleteButton';

type TableColumn<Entry> = {
  title: string;
  field: keyof Entry;
};

//TODO: 汎用的に使えるコンポーネントを抽出する
export type TableProps<Entry> = {
  data: Task[];
  columns: TableColumn<Task>[];
};

export const Table = <Entry extends Record<string, string | number>>({
  data,
  columns,
}: TableProps<Entry>) => {
  if (!data?.length) {
    return (
      <div className="flex flex-col items-center justify-center text-gray-500 bg-white h-80">
        <h4>No Entries Found</h4>
      </div>
    );
  }
  return (
    <div className="flex flex-col">
      <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <div className="overflow-hidden border-b border-gray-200 shadow sm:rounded-lg">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  {columns.map((column, index) => (
                    <th
                      key={column.title + index}
                      scope="col"
                      className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                    >
                      {column.title}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {data.map((entry, entryIndex) => (
                  <tr
                    key={entryIndex}
                    className="odd:bg-white even:bg-gray-100"
                  >
                    {columns.map(({ field, title }, columnIndex) => (
                      <td
                        key={title + columnIndex}
                        className="px-6 py-4 text-sm font-medium text-gray-900 whitespace-nowrap"
                      >
                        {entry[field] as string}
                        {/* TODO: アサーションをやめる */}
                      </td>
                    ))}
                    <td className="px-6 py-4 text-sm font-medium text-gray-900 whitespace-nowrap">
                      <FormDialog task={entry} />
                    </td>
                    <td>
                      <DeleteButton taskId={entry.taskId} />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};
