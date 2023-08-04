export type User = {
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
  uid:       string
  name:      string
  email:     string
  themes:    Theme[]
}

export type Theme = GormModel & {
  name:     string    
  ideas:    Idea[]
  elements: Element[]
  userUID:  string
}

export type Idea = GormModel & {
  name:    string
  theme_id: string
}

export type Element = GormModel & {
  name:    string
  theme_id: string
}

type GormModel = {
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
  ID:        string       
}