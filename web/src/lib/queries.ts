import {initialize} from "./scoretrak-queries";
import {scoretrakAxiosInstance} from "./api";

// @ts-ignore
export const scoretrak = initialize(scoretrakAxiosInstance, {})
