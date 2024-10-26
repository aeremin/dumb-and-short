import 'bootstrap/dist/css/bootstrap.min.css';
import {SuccessCard} from "@/app/created/[id]/success_card";

export function generateStaticParams() {
  return [{ id: "0" }];
}

export default async function Created({ params }: { params: Promise<{ id: string }>}) {
  const { id } = await params;
  return <SuccessCard  id={id}/>
}
