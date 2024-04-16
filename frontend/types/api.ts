export interface API {
  ID: number;
  name: string;
  path: string;
  method: string;
  handle_mode: HandleMode;
  body: string;
  javascript: string;
  replace: string;
  replace_with: string;
}

export enum HandleMode {
  ReplaceBody = "ReplaceBody",
  ModifyBody = "ModifyBody",
  JavaScript = "JavaScript",
}
