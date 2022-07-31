import { gql } from 'apollo-angular'
import { Injectable } from '@angular/core'
import * as Apollo from 'apollo-angular'
export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>
}
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string
  String: string
  Boolean: boolean
  Int: number
  Float: number
}

export type BasePaginationFilter = {
  count?: InputMaybe<Scalars['Int']>
  offset?: InputMaybe<Scalars['Int']>
}

export type DocumentReadable = File & {
  __typename?: 'DocumentReadable'
  blob: Scalars['String']
  contentType?: Maybe<Scalars['String']>
  extension?: Maybe<Scalars['String']>
  id: Scalars['ID']
  name?: Maybe<Scalars['String']>
  preview?: Maybe<Scalars['String']>
  size: Scalars['Int']
}

export type File = {
  contentType?: Maybe<Scalars['String']>
  extension?: Maybe<Scalars['String']>
  id: Scalars['ID']
  name?: Maybe<Scalars['String']>
}

export type FileLink = File & {
  __typename?: 'FileLink'
  contentType?: Maybe<Scalars['String']>
  extension?: Maybe<Scalars['String']>
  id: Scalars['ID']
  link: Scalars['String']
  name?: Maybe<Scalars['String']>
}

export type Mutation = {
  __typename?: 'Mutation'
  upsertProject?: Maybe<Project>
}

export type MutationUpsertProjectArgs = {
  input?: InputMaybe<ProjectInput>
}

export type Project = {
  __typename?: 'Project'
  description?: Maybe<Scalars['String']>
  id: Scalars['ID']
  key: Scalars['String']
  name: Scalars['String']
  technologies?: Maybe<Array<Technology>>
}

export type ProjectFilter = {
  count?: InputMaybe<Scalars['Int']>
  name?: InputMaybe<Scalars['String']>
  offset?: InputMaybe<Scalars['Int']>
  technologies?: InputMaybe<Array<InputMaybe<Scalars['String']>>>
}

export type ProjectInput = {
  description?: InputMaybe<Scalars['String']>
  id?: InputMaybe<Scalars['ID']>
  key?: InputMaybe<Scalars['String']>
  name?: InputMaybe<Scalars['String']>
  technologies?: InputMaybe<Array<InputMaybe<TechnologyInput>>>
}

export type Query = {
  __typename?: 'Query'
  projects: Array<Project>
}

export type QueryProjectsArgs = {
  filters?: InputMaybe<ProjectFilter>
}

export type Technology = {
  __typename?: 'Technology'
  description?: Maybe<Scalars['String']>
  id: Scalars['ID']
  key: Scalars['String']
  name: Scalars['String']
}

export type TechnologyInput = {
  description?: InputMaybe<Scalars['String']>
  id?: InputMaybe<Scalars['ID']>
  key?: InputMaybe<Scalars['String']>
  name?: InputMaybe<Scalars['String']>
}

export type GetAllProjectsQueryVariables = Exact<{
  filter?: InputMaybe<ProjectFilter>
}>

export type GetAllProjectsQuery = {
  __typename?: 'Query'
  projects: Array<{
    __typename?: 'Project'
    id: string
    key: string
    name: string
    description?: string | null
    technologies?: Array<{
      __typename?: 'Technology'
      id: string
      key: string
      name: string
      description?: string | null
    }> | null
  }>
}

export const GetAllProjectsDocument = gql`
  query GetAllProjects($filter: ProjectFilter) {
    projects(filters: $filter) {
      id
      key
      name
      description
      technologies {
        id
        key
        name
        description
      }
    }
  }
`

@Injectable({
  providedIn: 'root',
})
export class GetAllProjectsGQL extends Apollo.Query<
  GetAllProjectsQuery,
  GetAllProjectsQueryVariables
> {
  override document = GetAllProjectsDocument

  constructor(apollo: Apollo.Apollo) {
    super(apollo)
  }
}
