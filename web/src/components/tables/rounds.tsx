import {scoretrak} from "../../lib/queries";
import {createColumnHelper, flexRender, getCoreRowModel, useReactTable} from "@tanstack/react-table";
import {RoundList} from "../../lib/scoretrak-queries";

const columnHelper = createColumnHelper<RoundList>()

const columns = [
  columnHelper.accessor("id", {header: "ID"}),
  columnHelper.accessor("round_number", {}),
  columnHelper.accessor("note", {}),
  columnHelper.accessor("err", {}),
  columnHelper.accessor("started_at", {}),
  columnHelper.accessor("finished_at", {}),
  columnHelper.accessor("competition_id", {}),
  // columnHelper.accessor("create_time", {}),
  // columnHelper.accessor("update_time", {}),
  columnHelper.display({id: "edit", cell: props => <p onClick={() => console.log(props.row.original.id)}>Edit</p>})
]

export function RoundTable() {
  const { data, isLoading, isError } = scoretrak.queries.useListRound()
  const table = useReactTable({columns, data: data ?? [], getCoreRowModel: getCoreRowModel<RoundList>()})

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
