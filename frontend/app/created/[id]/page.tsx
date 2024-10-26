'use client'

import 'bootstrap/dist/css/bootstrap.min.css';
import Card from "react-bootstrap/Card";

export default async function Created({ params }: { params: Promise<{ id: string }>}) {
  const { id } = await params;
  return (<div>
    <Card body={false}>
      <Card.Header>Success!</Card.Header>
      <Card.Body>
        Link #{id} was successfully created!
      </Card.Body>
    </Card>
  </div>)
}
