'use client'

import 'bootstrap/dist/css/bootstrap.min.css';

import { useRouter } from 'next/router'

import Card from "react-bootstrap/Card";

export default function Created() {
  const router = useRouter()
  const id = router.query.id as string;
  return <Card body={false}>
    <Card.Header>Success!</Card.Header>
    <Card.Body>
      Link #{id} was successfully created!
    </Card.Body>
  </Card>
}
