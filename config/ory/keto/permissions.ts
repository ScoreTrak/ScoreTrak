import {Namespace, Context} from "@ory/keto-namespace-types"

class User implements Namespace {
}

class Competition implements Namespace {
  related: {
    owners: User[]
    staff: User[]
    competitors: Team[]
  }

  permits = {
    view: (ctx: Context): boolean =>
      this.related.viewers.includes(ctx.subject) ||
      this.related.editors.includes(ctx.subject) ||
      this.related.owners.includes(ctx.subject) ||
      this.related.parents.traverse((parent) => parent.permits.view(ctx)),
  }
}

class Team implements Namespace {
  related: {
    competitors: User[]
  }

  permits = {}
}

class Host implements Namespace {

}

class HostService implements Namespace {

}

class Round implements Namespace {

}

class Service implements Namespace {

}

class Property implements Namespace {
  related: {
    hostservice: HostService
  }
}
