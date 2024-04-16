import type { API } from "./api";

export interface Domain {
  ID: number;
  name: string;
  domain: string;
  desc: string;
  apis: API[];
  api_count: number;
}
