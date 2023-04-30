import {scoretrak} from "../../lib/queries";
import {createColumnHelper, flexRender, getCoreRowModel, useReactTable} from "@tanstack/react-table";
import {HostList, TeamList} from "../../lib/scoretrak-queries";
import {TeamHostsList} from "../../lib/scoretrak-queries";
import {useTeam} from "../../contexts/TeamContext";

const columnHelper = createColumnHelper<TeamHostsList>()

const columns = [
  columnHelper.accessor("id", {header: "ID"}),
  columnHelper.accessor("hidden", {}),
  columnHelper.accessor("pause", {}),
  columnHelper.accessor("address", {}),
  // columnHelper.accessor("address_list_range", {}),
  // columnHelper.accessor("editable", {}),
  columnHelper.accessor("team_id", {}),
  // columnHelper.accessor("create_time", {}),
  // columnHelper.accessor("update_time", {}),
  columnHelper.display({id: "edit", cell: props => <p onClick={() => console.log(props.row.original.id)}>Edit</p>})
]

export function HostTable() {
  const {teamId} = useTeam()
  const { data, isLoading, isError } = scoretrak.queries.useListTeamHosts(teamId)
  const table = useReactTable({columns, data: data ?? [], getCoreRowModel: getCoreRowModel<TeamHostsList>()})

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
