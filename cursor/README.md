# LBMesh Cursor Rules

The [`.cursorrules`](.cursorrules) file in this directory guides Cursor's AI to author and manage LBMesh recipes (`.lbmesh/recipe.yaml`) correctly.

## Usage

Copy `.cursorrules` into the root of your project, or reference its contents from your project's Cursor rules configuration.

## Auditing recipes

To audit a recipe, run:

```bash
lbmesh recipe lint <file>
```
