import 'bootstrap/dist/css/bootstrap.min.css';

import {redirect} from "next/navigation";

import {resolve} from "@/api"

export function generateStaticParams() {
  return [{ id: "0" }];
}

export default async function Redirect({ params }: { params: Promise<{ id: string }>}) {
  const { id } = await params;
  const url = await resolve(id);
  redirect(url);
}
