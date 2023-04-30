import {scoretrak} from "../../lib/queries";
import {useForm} from "react-hook-form";
import {queryKeys, TeamCreate} from "../../lib/scoretrak-queries";
import {queryClient} from "../../lib/query-client";


export function CreateTeamForm() {
  const {mutate} = scoretrak.mutations.useCreateTeam()
  const {handleSubmit, register, reset} = useForm<TeamCreate>({})
  const onSubmit = (data: TeamCreate) => {
    // @ts-ignore
    mutate(data, {
      onSuccess: () => {
        console.log(data)
        reset()
        return queryClient.invalidateQueries(queryKeys.listTeam())
      }
    })
  }

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <label>Name</label>
        <input {...register("name")} />
        <input {...register("display_name")}/>
        <input type={"number"} {...register("number")}/>
        <button type={"submit"}>Create Team</button>
      </form>
    </>
  )
}
