import {useForm} from "react-hook-form";
import {CompetitionCreate, CompetitionUpdate, queryKeys} from "../../lib/scoretrak-queries";
import {scoretrak} from "../../lib/queries";
import {queryClient} from "../../App";


export function CreateCompetitionForm() {
  const {mutate} = scoretrak.mutations.useCreateCompetition()
  const { handleSubmit, register, formState: { errors }, reset } = useForm<CompetitionCreate>()
  const onSubmit = (data: CompetitionCreate) => {
    mutate(data, {onSuccess: () => {
      console.log(data)
      reset()
      return queryClient.invalidateQueries(queryKeys.listCompetition())
      }})
  }

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <label>Name</label>
        <input {...register("name")} />
        <label>Display Name</label>
        <input {...register("display_name")} />
        <label>Round Duration</label>
        <input type={"number"} {...register("round_duration")} defaultValue={60}/>
        {/*<input type={"checkbox"} {...register("pause")} />*/}
        {/*<input type={"checkbox"} {...register("hidden")} />*/}
        <button type="submit">Create Competition</button>
      </form>
    </>
  )
}
