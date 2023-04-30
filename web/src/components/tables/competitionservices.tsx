import {scoretrak} from "../../lib/queries";
import {createColumnHelper, flexRender, getCoreRowModel, useReactTable} from "@tanstack/react-table";
import {ServiceList} from "../../lib/scoretrak-queries";
import {CompetitionServicesList} from "../../lib/scoretrak-queries";
import {useCompetition} from "../../contexts/CompetitionContext";

const columnHelper = createColumnHelper<CompetitionServicesList>()

const columns = [
  columnHelper.accessor("id", {header: "ID"}),
  columnHelper.accessor("name", {header: "Name"}),
  columnHelper.accessor("display_name", {header: "Display Name"}),
  columnHelper.accessor("hidden", {}),
  columnHelper.accessor("pause", {}),
  columnHelper.accessor("competition_id", {}),
  // columnHelper.accessor("create_time", {}),
  // columnHelper.accessor("update_time", {}),
  columnHelper.display({id: "edit", cell: props => <p onClick={() => console.log(props.row.original.id)}>Edit</p>})
]

export function CompetitionServiceTable() {
  const {competitionId} = useCompetition()
  const { data, isLoading, isError } = scoretrak.queries.useListCompetitionServices(competitionId)
  const table = useReactTable({columns, data: data ?? [], getCoreRowModel: getCoreRowModel<CompetitionServicesList>()})

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
