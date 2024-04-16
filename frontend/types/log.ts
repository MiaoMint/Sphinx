export interface Log {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  status: number;
  latency: number;
  ip: string;
  method: string;
  path: string;
  domain: string;
  api_id: number;
}
