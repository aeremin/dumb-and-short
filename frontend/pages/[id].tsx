import {resolve} from "@/api"
import {useRouter} from "next/router";
import {useEffect} from "react";

export default function Redirect() {
  const router = useRouter()
  const id = router.query.id as string;

  useEffect(() => {
    resolve(id).then(url => router.push(url));
  }, [id])

  return <div/>
}
