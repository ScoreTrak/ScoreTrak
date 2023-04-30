import {scoretrak} from "../../lib/queries";
import {createColumnHelper, flexRender, getCoreRowModel, useReactTable} from "@tanstack/react-table";
import {TeamList} from "../../lib/scoretrak-queries";

const columnHelper = createColumnHelper<TeamList>()

const columns = [
  columnHelper.accessor("id", {header: "ID"}),
  columnHelper.accessor("number", {}),
  columnHelper.accessor("name", {header: "Name"}),
  columnHelper.accessor("display_name", {header: "Display Name"}),
  columnHelper.accessor("hidden", {}),
  columnHelper.accessor("pause", {}),
  // columnHelper.accessor("create_time", {}),
  // columnHelper.accessor("update_time", {}),
  columnHelper.display({id: "edit", cell: props => <p onClick={() => console.log(props.row.original.id)}>Edit</p>})
]

export function TeamTable() {
  const { data, isLoading, isError } = scoretrak.queries.useListTeam()
  const table = useReactTable({columns, data: data ?? [], getCoreRowModel: getCoreRowModel<TeamList>()})

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
