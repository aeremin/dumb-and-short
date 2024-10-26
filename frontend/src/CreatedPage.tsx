import 'bootstrap/dist/css/bootstrap.min.css';

import Card from "react-bootstrap/Card";
import {useLoaderData} from "react-router-dom";

export default function CreatedPage() {
  const id = useLoaderData() as string;

  return <Card body={false}>
    <Card.Header>Success!</Card.Header>
    <Card.Body>
      Link #{id} was successfully created!
    </Card.Body>
  </Card>
}
