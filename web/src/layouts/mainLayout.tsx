import {PropsWithChildren} from "react";
import Nav from "../components/Nav";


export function MainLayout({children}: PropsWithChildren) {
  return (
    <>
      <Nav />
      {children}
    </>
  )
}
