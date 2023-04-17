import {scoretrak} from "../../lib/queries";
import {useForm} from "react-hook-form";
import {queryKeys, TeamCompetitionRead, TeamCreate} from "../../lib/scoretrak-queries";
import {queryClient} from "../../App";


export function CreateCompetitionTeamForm({ competitionId }: { competitionId: number }) {
  const {mutate} = scoretrak.mutations.useCreateTeam()
  const {handleSubmit, register, reset} = useForm<TeamCreate>({
    defaultValues: {
      competition_id: competitionId
    }
  })
  const onSubmit = (data: TeamCreate) => {
    mutate(data, {
      onSuccess: () => {
        console.log(data)
        reset()
        return queryClient.invalidateQueries(queryKeys.listCompetitionTeams(competitionId))
      }
    })
  }

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <label>Name</label>
        <input {...register("name")} />
        <input type="hidden" {...register("competition_id")}/>
        <button type={"submit"}>Create Team</button>
      </form>
    </>
  )
}
