import {createBrowserRouter} from "react-router-dom";
import {RootPage} from "./routes/root";
import {ErrorPage} from "./routes/error";
import {AuthLoginPage} from "./routes/auth/login";
import {AuthRegisterPage} from "./routes/auth/register";
import {TeamsPage} from "./routes/dashboard/teams";
import {AuthErrorPage} from "./routes/auth/error";
import {ServicesPage} from "./routes/dashboard/services";
import {HostServicesPage} from "./routes/dashboard/hostservices";
import {PropertiesPage} from "./routes/dashboard/properties";
import {ReportsPage} from "./routes/dashboard/reports";
import {RoundsPage} from "./routes/dashboard/rounds";
import {ChecksPage} from "./routes/dashboard/checks";
import {IndexPage} from "./routes";
import {CompetitionPortalPage} from "./routes/c/competition";
import {CompetitionProvider} from "./contexts/CompetitionContext";
import {CompetitionsPage} from "./routes/dashboard/competitions";
import {DashboardPage} from "./routes/dashboard";
import {HostsPage} from "./routes/dashboard/hosts";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootPage />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/",
        element: <IndexPage />,
        index: true
      },
      {
        path: "dashboard",
        element: <DashboardPage />,
        children: [
          {
            path: "competitions",
            element: <CompetitionsPage />
          },
          {
            path: "teams",
            element: <TeamsPage />
          },
          {
            path: "hosts",
            element: <HostsPage />
          },
          {
            path: "hostservices",
            element: <HostServicesPage />
          },
          {
            path: "services",
            element: <ServicesPage />
          },
          {
            path: "properties",
            element: <PropertiesPage />
          },
          {
            path: "reports",
            element: <ReportsPage />
          },
          {
            path: "rounds",
            element: <RoundsPage />
          },
          {
            path: "checks",
            element: <ChecksPage />
          },
        ]
      },
          {
            path: "c/:competitionId",
            element: <CompetitionProvider><CompetitionPortalPage /></CompetitionProvider>,
            children: [
              {
                path: "scoreboard",
              },
              {
                path: "ranks",
              },
              {
                path: "health",
              },
            ]
          },
      {
        path: "auth",
        children: [
          {
            path: "login",
            element: <AuthLoginPage />
          },
          {
            path: "register",
            element: <AuthRegisterPage />
          },
          {
            path: "error",
            element: <AuthErrorPage />
          }
        ]
      },
    ]
  }
])
