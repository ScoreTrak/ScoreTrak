import {scoretrak} from "../../lib/queries";
import {createColumnHelper, flexRender, getCoreRowModel, useReactTable} from "@tanstack/react-table";
import {CheckList, CompetitionList} from "../../lib/scoretrak-queries";

const columnHelper = createColumnHelper<CompetitionList>()

const columns = [
  columnHelper.accessor("id", {header: "ID"}),
  columnHelper.accessor("name", {header: "ID"}),
  columnHelper.accessor("display_name", {header: "ID"}),
  columnHelper.accessor("hidden", {}),
  columnHelper.accessor("pause", {}),
  // columnHelper.accessor("create_time", {}),
  // columnHelper.accessor("update_time", {}),
  columnHelper.display({id: "edit", cell: props => <p onClick={() => console.log(props.row.original.id)}>Edit</p>}),
  columnHelper.display({id: "start", cell: props => <p onClick={() => console.log(props.row.original.id)}>Start</p>}),
  columnHelper.display({id: "pause", cell: props => <p onClick={() => console.log(props.row.original.id)}>Pause</p>})
]

export function CompetitionTable() {
  const { data, isLoading, isError } = scoretrak.queries.useListCompetition()
  const table = useReactTable({columns, data: data ?? [], getCoreRowModel: getCoreRowModel<CompetitionList>()})

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
