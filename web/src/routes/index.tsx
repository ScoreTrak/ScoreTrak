import {scoretrak} from "../lib/queries";
import {Link} from "react-router-dom";


export function IndexPage() {
  const {data} = scoretrak.queries.useListCompetition()
  return (
    <>
        {
          data ?
          <>
      <h2>Available Competitions</h2>
      <div className="grid grid-cols-4">
          {data.map((comp, idx) => {
            return (
              <Link to={`/c/${comp.id}`} key={idx}>
                <div key={idx} className={"p-4"}>
                  <h3 className={"text-lg"}>{comp.name}</h3>
                </div>
              </Link>
            )
          })}
      </div>
          </> :
            <h2>No Available Competitions</h2>
        }

    </>
  )
}
