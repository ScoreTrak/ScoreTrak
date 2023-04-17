import { Route } from "@tanstack/router";
import { indexRoute } from "..";
import { rootRoute } from "../root";

import {z} from "zod"
import {baseLayout} from "../layouts/baseLayout";
import {unauthenticatedLayout} from "../layouts/unauthenticatedLayout";

export const authSearchSchema = z.object({
  flow: z.string().uuid().nullable().optional(),
  aal1: z.string().nullable().optional(),
  aal2: z.string().nullable().optional(),
})

type AuthSearch = z.infer<typeof authSearchSchema>

export const authRoute = new Route({
  getParentRoute: () => unauthenticatedLayout,
  path: "auth",
  // validateSearch: authSearchSchema
})
