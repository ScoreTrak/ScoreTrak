import {scoretrak} from "../../lib/queries";
import {createColumnHelper, flexRender, getCoreRowModel, useReactTable} from "@tanstack/react-table";
import {CheckList} from "../../lib/scoretrak-queries";

const columnHelper = createColumnHelper<CheckList>()

const columns = [
  columnHelper.accessor("id", {header: "ID"}),
  columnHelper.accessor("log", {}),
  columnHelper.accessor("error", {}),
  columnHelper.accessor("passed", {}),
  columnHelper.accessor("hidden", {}),
  columnHelper.accessor("pause", {}),
  columnHelper.accessor("host_service_id", {}),
  columnHelper.accessor("round_id", {}),
  columnHelper.accessor("team_id", {}),
  // columnHelper.accessor("create_time", {}),
  // columnHelper.accessor("update_time", {}),
  columnHelper.display({id: "edit", cell: props => <p onClick={() => console.log(props.row.original.id)}>Edit</p>})
]

export function CheckTable() {
  const { data, isLoading, isError } = scoretrak.queries.useListCheck()
  const table = useReactTable({columns, data: data ?? [], getCoreRowModel: getCoreRowModel<CheckList>()})

  return (
    <>
      <table>
        <thead>
        {table.getHeaderGroups().map(headerGroup => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map(header => (
              <th key={header.id}>
                {header.isPlaceholder
                  ? null
                  : flexRender(
                    header.column.columnDef.header,
                    header.getContext()
                  )}
              </th>
            ))}
          </tr>
        ))}
        </thead>
        <tbody>
        {table.getRowModel().rows.map(row => (
          <tr key={row.id}>
            {row.getVisibleCells().map(cell => (
              <td key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </td>
            ))}
          </tr>
        ))}
        </tbody>
      </table>
    </>
  )
}
