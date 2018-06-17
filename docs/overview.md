[//]: # (This is generated file, please don't edit it yourself.)

## Overview

Go source code linter that brings checks that are currently not implemented in other linters.

## Checkers:

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr>
      <tr>
        <td><a href="#appendCombine-ref">appendCombine</a></td>
        <td>Detects `append` chains to the same slice that can be done in a single `append` call.</td>
      </tr>
      <tr>
        <td><a href="#builtinShadow-ref">builtinShadow</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#captLocal-ref">captLocal</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#docStub-ref">docStub</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#elseif-ref">elseif</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#flagDeref-ref">flagDeref</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#paramTypeCombine-ref">paramTypeCombine</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#ptrToRefParam-ref">ptrToRefParam</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#rangeExprCopy-ref">rangeExprCopy</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#rangeValCopy-ref">rangeValCopy</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#singleCaseSwitch-ref">singleCaseSwitch</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#stdExpr-ref">stdExpr</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#switchTrue-ref">switchTrue</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#typeSwitchVar-ref">typeSwitchVar</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#typeUnparen-ref">typeUnparen</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#underef-ref">underef</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#unexportedCall-ref">unexportedCall</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#unnamedResult-ref">unnamedResult</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#unslice-ref">unslice</a></td>
        <td></td>
      </tr>
</table>

**Experimental:**

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr>
      <tr>
        <td><a href="#longChain-ref">longChain</a></td>
        <td></td>
      </tr>
      <tr>
        <td><a href="#unusedParam-ref">unusedParam</a></td>
        <td></td>
      </tr>
</table>



<a name="appendCombine-ref"></a>
## appendCombine
Detects `append` chains to the same slice that can be done in a single `append` call.


**Before:**
```go
xs = append(xs, 1)
xs = append(xs, 2)

```

**After:**
```go
xs = append(xs, 1, 2)

```


<a name="builtinShadow-ref"></a>
## builtinShadow


**Before:**
```go

```

**After:**
```go

```

`builtinShadow` is syntax-only checker (fast).
<a name="captLocal-ref"></a>
## captLocal


**Before:**
```go

```

**After:**
```go

```

`captLocal` is syntax-only checker (fast).
<a name="docStub-ref"></a>
## docStub


**Before:**
```go

```

**After:**
```go

```

`docStub` is syntax-only checker (fast).
<a name="elseif-ref"></a>
## elseif


**Before:**
```go

```

**After:**
```go

```

`elseif` is syntax-only checker (fast).
<a name="flagDeref-ref"></a>
## flagDeref


**Before:**
```go

```

**After:**
```go

```

`flagDeref` is syntax-only checker (fast).
<a name="paramTypeCombine-ref"></a>
## paramTypeCombine


**Before:**
```go

```

**After:**
```go

```

`paramTypeCombine` is syntax-only checker (fast).
<a name="ptrToRefParam-ref"></a>
## ptrToRefParam


**Before:**
```go

```

**After:**
```go

```


<a name="rangeExprCopy-ref"></a>
## rangeExprCopy


**Before:**
```go

```

**After:**
```go

```


<a name="rangeValCopy-ref"></a>
## rangeValCopy


**Before:**
```go

```

**After:**
```go

```


<a name="singleCaseSwitch-ref"></a>
## singleCaseSwitch


**Before:**
```go

```

**After:**
```go

```

`singleCaseSwitch` is syntax-only checker (fast).
<a name="stdExpr-ref"></a>
## stdExpr


**Before:**
```go

```

**After:**
```go

```


<a name="switchTrue-ref"></a>
## switchTrue


**Before:**
```go

```

**After:**
```go

```

`switchTrue` is syntax-only checker (fast).
<a name="typeSwitchVar-ref"></a>
## typeSwitchVar


**Before:**
```go

```

**After:**
```go

```


<a name="typeUnparen-ref"></a>
## typeUnparen


**Before:**
```go

```

**After:**
```go

```

`typeUnparen` is syntax-only checker (fast).
<a name="underef-ref"></a>
## underef


**Before:**
```go

```

**After:**
```go

```


<a name="unexportedCall-ref"></a>
## unexportedCall


**Before:**
```go

```

**After:**
```go

```

`unexportedCall` is very opinionated.
<a name="unnamedResult-ref"></a>
## unnamedResult


**Before:**
```go

```

**After:**
```go

```


<a name="unslice-ref"></a>
## unslice


**Before:**
```go

```

**After:**
```go

```


<a name="longChain-ref"></a>
## longChain


**Before:**
```go

```

**After:**
```go

```


<a name="unusedParam-ref"></a>
## unusedParam


**Before:**
```go

```

**After:**
```go

```

