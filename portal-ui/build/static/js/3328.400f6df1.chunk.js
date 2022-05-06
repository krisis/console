"use strict";(self.webpackChunkportal_ui=self.webpackChunkportal_ui||[]).push([[3328],{63328:function(e,n,t){t.r(n),t.d(n,{default:function(){return q}});var s=t(29439),i=t(1413),o=t(72791),r=t(60364),a=t(11135),l=t(25787),c=t(61889),u=t(64554),d=t(26181),f=t.n(d),x=t(42649),p=t(23814),h=t(81207),m=t(42209),Z=t(56087),v=t(38442),g=t(75578),j=t(45902),b=t(80184),E=function(e){var n=e.icon,t=void 0===n?null:n,s=e.label,i=void 0===s?null:s;return(0,b.jsxs)(u.Z,{sx:{display:"flex"},children:[(0,b.jsx)("div",{style:{height:16,width:16,display:"flex",alignItems:"center",marginTop:5},children:t}),(0,b.jsx)("div",{style:{marginLeft:t?5:"none"},children:i})]})},k=t(56375),T=t(45987),S=t(36151),C=t(72401),y=["isLoading","onClick","label"],_=function(e){var n=e.isLoading,t=e.onClick,s=e.label,o=(0,T.Z)(e,y);return(0,b.jsx)(S.Z,(0,i.Z)((0,i.Z)({color:"primary",onClick:t,variant:"text",sx:{padding:0,margin:0,alignItems:"flex-start",justifyContent:"flex-start",display:"inline-flex",height:"auto",textDecoration:"underline",color:"#2781B0","&:hover":{background:"#ffffff",textDecoration:"underline"}},disableRipple:!0,disableFocusRipple:!0},o),{},{children:n?(0,b.jsx)(C.Z,{style:{width:16,height:16}}):s}))},N=t(13400),I=t(33548),O=["disabled","onClick"],R=(0,l.Z)((function(e){return(0,a.Z)({root:{"&:hover":{backgroundColor:"#E2E2E2"}}})}))((function(e){var n=e.disabled,t=e.onClick,s=(0,T.Z)(e,O);return(0,b.jsx)(N.Z,(0,i.Z)((0,i.Z)({size:"small",disabled:n,onClick:t},s),{},{children:(0,b.jsx)(I.Z,{})}))})),F=function(e){var n=e.resourceName,t=e.iamScopes,s=e.secureCmpProps,o=void 0===s?{}:s,r=e.children;return(0,b.jsx)(v.s,(0,i.Z)((0,i.Z)({scopes:t,resource:n,errorProps:{disabled:!0}},o),{},{children:r}))},P=function(e){var n=e.isLoading,t=void 0===n||n,s=e.resourceName,i=void 0===s?"":s,o=e.iamScopes,r=e.secureCmpProps,a=void 0===r?{}:r,l=e.property,c=void 0===l?null:l,d=e.value,f=void 0===d?null:d,x=e.onEdit;return(0,b.jsxs)(u.Z,{sx:{display:"flex",alignItems:"baseline",justifyContent:"flex-start"},children:[(0,b.jsx)(j.Z,{label:c,value:(0,b.jsx)(F,{resourceName:i,iamScopes:o,secureCmpProps:a,children:(0,b.jsx)(_,{isLoading:t,onClick:x,label:f})})}),(0,b.jsx)(F,{resourceName:i,iamScopes:o,secureCmpProps:a,children:(0,b.jsx)(R,{onClick:x,sx:{background:"#f8f8f8",marginLeft:"3px",top:3,"& .min-icon":{width:"16px",height:"16px"}}})})]})},B=t(45248),U=function(e){var n=e.bucketSize;return(0,b.jsxs)(u.Z,{sx:{display:"flex",alignItems:"center","& .min-icon":{height:37,width:37}},children:[(0,b.jsx)(k.Mhr,{}),(0,b.jsxs)(u.Z,{sx:{display:"flex",alignItems:"flex-start",justifyContent:"center",flexFlow:"column",marginLeft:"20px",fontSize:"19px"},children:[(0,b.jsx)("label",{style:{fontWeight:600},children:"Reported Usage:"}),(0,b.jsx)("label",{children:(0,B.ae)(n)})]})]})},G=function(e){var n=e.quota;return(0,b.jsxs)(u.Z,{sx:{display:"flex",alignItems:"center","& .min-icon":{height:37,width:37}},children:[(0,b.jsx)(k.sjJ,{}),(0,b.jsxs)(u.Z,{sx:{display:"flex",alignItems:"flex-start",justifyContent:"center",flexFlow:"column",marginLeft:"20px",fontSize:"19px"},children:[(0,b.jsxs)("label",{style:{fontWeight:600,textTransform:"capitalize"},children:[null===n||void 0===n?void 0:n.type," Quota"]}),(0,b.jsxs)("label",{children:[" ",(0,B.ae)("".concat(null===n||void 0===n?void 0:n.quota),!0)]})]})]})},A=t(50896),L=(0,g.Z)(o.lazy((function(){return Promise.all([t.e(5609),t.e(3631)]).then(t.bind(t,33690))}))),w=(0,g.Z)(o.lazy((function(){return t.e(1604).then(t.bind(t,1604))}))),M=(0,g.Z)(o.lazy((function(){return t.e(8391).then(t.bind(t,18391))}))),D=(0,g.Z)(o.lazy((function(){return t.e(402).then(t.bind(t,70402))}))),K=(0,g.Z)(o.lazy((function(){return Promise.all([t.e(1918),t.e(1705)]).then(t.bind(t,71705))}))),z=(0,g.Z)(o.lazy((function(){return t.e(1581).then(t.bind(t,1581))}))),V={display:"grid",gridTemplateColumns:{xs:"1fr",sm:"2fr 1fr"},gridAutoFlow:{xs:"dense",sm:"row"},gap:2},Y=(0,r.$j)((function(e){return{session:e.console.session,distributedSetup:e.system.distributedSetup,loadingBucket:e.buckets.bucketDetails.loadingBucket,bucketInfo:e.buckets.bucketDetails.bucketInfo}}),{setErrorSnackMessage:x.Ih,setBucketDetailsLoad:m.d5}),q=(0,l.Z)((function(e){return(0,a.Z)((0,i.Z)((0,i.Z)({},p.bK),p.VI))}))(Y((function(e){var n=e.classes,t=e.match,r=e.distributedSetup,a=e.setErrorSnackMessage,l=e.loadingBucket,d=e.bucketInfo,x=e.setBucketDetailsLoad,p=(0,o.useState)(null),m=(0,s.Z)(p,2),g=m[0],T=m[1],S=(0,o.useState)("0"),C=(0,s.Z)(S,2),y=C[0],_=C[1],N=(0,o.useState)(!1),I=(0,s.Z)(N,2),O=I[0],R=I[1],F=(0,o.useState)(!1),B=(0,s.Z)(F,2),Y=B[0],q=B[1],Q=(0,o.useState)(!1),J=(0,s.Z)(Q,2),W=J[0],$=J[1],H=(0,o.useState)(!0),X=(0,s.Z)(H,2),ee=X[0],ne=X[1],te=(0,o.useState)(!0),se=(0,s.Z)(te,2),ie=se[0],oe=se[1],re=(0,o.useState)(!0),ae=(0,s.Z)(re,2),le=ae[0],ce=ae[1],ue=(0,o.useState)(!0),de=(0,s.Z)(ue,2),fe=de[0],xe=de[1],pe=(0,o.useState)(!0),he=(0,s.Z)(pe,2),me=he[0],Ze=he[1],ve=(0,o.useState)(!0),ge=(0,s.Z)(ve,2),je=ge[0],be=ge[1],Ee=(0,o.useState)(!0),ke=(0,s.Z)(Ee,2),Te=ke[0],Se=ke[1],Ce=(0,o.useState)(!0),ye=(0,s.Z)(Ce,2),_e=ye[0],Ne=ye[1],Ie=(0,o.useState)(!1),Oe=(0,s.Z)(Ie,2),Re=Oe[0],Fe=Oe[1],Pe=(0,o.useState)(!1),Be=(0,s.Z)(Pe,2),Ue=Be[0],Ge=Be[1],Ae=(0,o.useState)(null),Le=(0,s.Z)(Ae,2),we=Le[0],Me=Le[1],De=(0,o.useState)(!1),Ke=(0,s.Z)(De,2),ze=Ke[0],Ve=Ke[1],Ye=(0,o.useState)(!1),qe=(0,s.Z)(Ye,2),Qe=qe[0],Je=qe[1],We=(0,o.useState)(null),$e=(0,s.Z)(We,2),He=$e[0],Xe=$e[1],en=(0,o.useState)(!1),nn=(0,s.Z)(en,2),tn=nn[0],sn=nn[1],on=(0,o.useState)(!1),rn=(0,s.Z)(on,2),an=rn[0],ln=rn[1],cn=(0,o.useState)(!1),un=(0,s.Z)(cn,2),dn=un[0],fn=un[1],xn=(0,o.useState)(!1),pn=(0,s.Z)(xn,2),hn=pn[0],mn=pn[1],Zn=t.params.bucketName,vn="n/a",gn="";null!==d&&(vn=d.access,gn=d.definition);var jn=(0,v.F)(Zn,[Z.Ft.S3_GET_BUCKET_OBJECT_LOCK_CONFIGURATION]),bn=(0,v.F)(Zn,[Z.Ft.S3_GET_BUCKET_ENCRYPTION_CONFIGURATION]),En=(0,v.F)(Zn,[Z.Ft.ADMIN_GET_BUCKET_QUOTA]);(0,o.useEffect)((function(){ce(!!l)}),[l,ce]),(0,o.useEffect)((function(){fe&&(bn?h.Z.invoke("GET","/api/v1/buckets/".concat(Zn,"/encryption/info")).then((function(e){e.algorithm&&(Ve(!0),T(e)),xe(!1)})).catch((function(e){"The server side encryption configuration was not found"===e.errorMessage&&(Ve(!1),T(null)),xe(!1)})):(Ve(!1),T(null),xe(!1)))}),[fe,Zn,bn]),(0,o.useEffect)((function(){me&&r&&h.Z.invoke("GET","/api/v1/buckets/".concat(Zn,"/versioning")).then((function(e){Fe(e.is_versioned),Ze(!1)})).catch((function(e){a(e),Ze(!1)}))}),[me,a,Zn,r]),(0,o.useEffect)((function(){je&&r&&(En?h.Z.invoke("GET","/api/v1/buckets/".concat(Zn,"/quota")).then((function(e){Me(e),e.quota?Ge(!0):Ge(!1),be(!1)})).catch((function(e){a(e),Ge(!1),be(!1)})):(Ge(!1),be(!1)))}),[je,Ze,a,Zn,r,En]),(0,o.useEffect)((function(){me&&r&&(jn?h.Z.invoke("GET","/api/v1/buckets/".concat(Zn,"/object-locking")).then((function(e){R(e.object_locking_enabled),ne(!1)})).catch((function(e){a(e),ne(!1)})):ne(!1))}),[ee,a,Zn,me,r,jn]),(0,o.useEffect)((function(){ie&&h.Z.invoke("GET","/api/v1/buckets").then((function(e){var n=f()(e,"buckets",[]).find((function(e){return e.name===Zn})),t=f()(n,"size","0");oe(!1),_(t)})).catch((function(e){oe(!1),a(e)}))}),[ie,a,Zn]),(0,o.useEffect)((function(){Te&&r&&h.Z.invoke("GET","/api/v1/buckets/".concat(Zn,"/replication")).then((function(e){var n=e.rules?e.rules:[];$(n.length>0),Se(!1)})).catch((function(e){a(e),Se(!1)}))}),[Te,a,Zn,r]),(0,o.useEffect)((function(){_e&&O&&h.Z.invoke("GET","/api/v1/buckets/".concat(Zn,"/retention")).then((function(e){Ne(!1),Je(!0),Xe(e)})).catch((function(e){Je(!1),Ne(!1),Xe(null)}))}),[_e,O,Zn]);var kn=function(){x(!0),ce(!0),oe(!0),Ze(!0),xe(!0),Ne(!0)};return(0,b.jsxs)(o.Fragment,{children:[an&&(0,b.jsx)(M,{open:an,selectedBucket:Zn,encryptionEnabled:ze,encryptionCfg:g,closeModalAndRefresh:function(){ln(!1),xe(!0)}}),dn&&(0,b.jsx)(z,{open:dn,selectedBucket:Zn,enabled:Ue,cfg:we,closeModalAndRefresh:function(){fn(!1),be(!0)}}),Y&&(0,b.jsx)(L,{bucketName:Zn,open:Y,actualPolicy:vn,actualDefinition:gn,closeModalAndRefresh:function(){q(!1),kn()}}),tn&&(0,b.jsx)(w,{bucketName:Zn,open:tn,closeModalAndRefresh:function(){sn(!1),kn()}}),hn&&(0,b.jsx)(D,{closeVersioningModalAndRefresh:function(e){mn(!1),e&&kn()},modalOpen:hn,selectedBucket:Zn,versioningCurrentState:Re}),(0,b.jsx)(A.Z,{children:"Summary"}),(0,b.jsxs)(c.ZP,{container:!0,spacing:1,children:[(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_BUCKET_POLICY],resource:Zn,children:(0,b.jsx)(c.ZP,{item:!0,xs:12,children:(0,b.jsxs)(u.Z,{sx:(0,i.Z)({},V),children:[(0,b.jsxs)(u.Z,{sx:(0,i.Z)({},V),children:[(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_BUCKET_POLICY],resource:Zn,children:(0,b.jsx)(P,{iamScopes:[Z.Ft.S3_PUT_BUCKET_POLICY],resourceName:Zn,property:"Access Policy:",value:vn.toLowerCase(),onEdit:function(){q(!0)},isLoading:le})}),(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_BUCKET_ENCRYPTION_CONFIGURATION],resource:Zn,children:(0,b.jsx)(P,{iamScopes:[Z.Ft.S3_PUT_BUCKET_ENCRYPTION_CONFIGURATION],resourceName:Zn,property:"Encryption:",value:ze?"Enabled":"Disabled",onEdit:function(){ln(!0)},isLoading:fe})}),(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_REPLICATION_CONFIGURATION],resource:Zn,children:(0,b.jsx)(j.Z,{label:"Replication:",value:(0,b.jsx)(E,{icon:W?(0,b.jsx)(k.E31,{}):(0,b.jsx)(k.dRf,{}),label:(0,b.jsx)("label",{className:n.textMuted,children:W?"Enabled":"Disabled"})})})}),(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_BUCKET_OBJECT_LOCK_CONFIGURATION],resource:Zn,children:(0,b.jsx)(j.Z,{label:"Object Locking:",value:(0,b.jsx)(E,{icon:O?(0,b.jsx)(k.E31,{}):(0,b.jsx)(k.dRf,{}),label:(0,b.jsx)("label",{className:n.textMuted,children:O?"Enabled":"Disabled"})})})}),(0,b.jsx)(u.Z,{className:n.spacerTop,children:(0,b.jsx)(j.Z,{label:"Tags:",value:(0,b.jsx)(K,{setErrorSnackMessage:a,bucketName:Zn})})}),(0,b.jsx)(P,{iamScopes:[Z.Ft.ADMIN_SET_BUCKET_QUOTA],resourceName:Zn,property:"Quota:",value:Ue?"Enabled":"Disabled",onEdit:function(){fn(!0)},isLoading:je})]}),(0,b.jsxs)(u.Z,{sx:{display:"grid",gridTemplateColumns:"1fr",alignItems:"flex-start"},children:[(0,b.jsx)(U,{bucketSize:y}),Ue&&we?(0,b.jsx)(G,{quota:we}):null]})]})})}),r&&(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_BUCKET_VERSIONING],resource:Zn,children:(0,b.jsxs)(c.ZP,{item:!0,xs:12,children:[(0,b.jsx)(A.Z,{children:"Versioning"}),(0,b.jsx)(u.Z,{sx:(0,i.Z)({},V),children:(0,b.jsx)(u.Z,{sx:(0,i.Z)({},V),children:(0,b.jsx)(P,{iamScopes:[Z.Ft.S3_PUT_BUCKET_VERSIONING],resourceName:Zn,property:"Current Status:",value:Re?"Versioned":"Unversioned (Default)",onEdit:function(){mn(!0)},isLoading:me})})})]})}),O&&(0,b.jsx)(v.s,{scopes:[Z.Ft.S3_GET_OBJECT_RETENTION],resource:Zn,children:(0,b.jsxs)(c.ZP,{item:!0,xs:12,children:[(0,b.jsx)(A.Z,{children:"Retention"}),(0,b.jsxs)(u.Z,{sx:{display:"grid",gridTemplateColumns:{xs:"1fr",sm:"2fr 1fr"},gridAutoFlow:{xs:"dense",sm:"row"},gap:2},children:[(0,b.jsxs)(u.Z,{sx:{display:"grid",gridTemplateColumns:{xs:"1fr",sm:"2fr 1fr"},gridAutoFlow:{xs:"dense",sm:"row"},gap:2},children:[(0,b.jsx)(P,{iamScopes:[Z.Ft.ADMIN_SET_BUCKET_QUOTA],resourceName:Zn,property:"Retention:",value:Qe?"Enabled":"Disabled",onEdit:function(){sn(!0)},isLoading:_e}),(0,b.jsx)(j.Z,{label:"Mode:",value:(0,b.jsx)("label",{className:n.textMuted,style:{textTransform:"capitalize"},children:He&&He.mode?He.mode:"-"})}),(0,b.jsx)(j.Z,{label:"Validity:",value:(0,b.jsxs)("label",{className:n.textMuted,style:{textTransform:"capitalize"},children:[He&&He.validity," ",He&&(1===He.validity?He.unit.slice(0,-1):He.unit)]})})]}),(0,b.jsx)(u.Z,{sx:{display:"grid",gridTemplateColumns:"1fr",alignItems:"flex-start"}})]})]})})]})]})})))},45902:function(e,n,t){var s=t(1413),i=(t(72791),t(53767)),o=t(80184);n.Z=function(e){var n=e.label,t=void 0===n?null:n,r=e.value,a=void 0===r?"-":r,l=e.orientation,c=void 0===l?"column":l,u=e.stkProps,d=void 0===u?{}:u,f=e.lblProps,x=void 0===f?{}:f,p=e.valProps,h=void 0===p?{}:p;return(0,o.jsxs)(i.Z,(0,s.Z)((0,s.Z)({direction:{xs:"column",sm:c}},d),{},{children:[(0,o.jsx)("label",(0,s.Z)((0,s.Z)({style:{marginRight:5,fontWeight:600}},x),{},{children:t})),(0,o.jsx)("label",(0,s.Z)((0,s.Z)({style:{marginRight:5,fontWeight:500}},h),{},{children:a}))]}))}},53767:function(e,n,t){var s=t(4942),i=t(63366),o=t(87462),r=t(72791),a=t(51184),l=t(45682),c=t(78519),u=t(82466),d=t(47630),f=t(93736),x=t(80184),p=["component","direction","spacing","divider","children"];function h(e,n){var t=r.Children.toArray(e).filter(Boolean);return t.reduce((function(e,s,i){return e.push(s),i<t.length-1&&e.push(r.cloneElement(n,{key:"separator-".concat(i)})),e}),[])}var m=(0,d.ZP)("div",{name:"MuiStack",slot:"Root",overridesResolver:function(e,n){return[n.root]}})((function(e){var n=e.ownerState,t=e.theme,i=(0,o.Z)({display:"flex"},(0,a.k9)({theme:t},(0,a.P$)({values:n.direction,breakpoints:t.breakpoints.values}),(function(e){return{flexDirection:e}})));if(n.spacing){var r=(0,l.hB)(t),c=Object.keys(t.breakpoints.values).reduce((function(e,t){return null==n.spacing[t]&&null==n.direction[t]||(e[t]=!0),e}),{}),d=(0,a.P$)({values:n.direction,base:c}),f=(0,a.P$)({values:n.spacing,base:c});i=(0,u.Z)(i,(0,a.k9)({theme:t},f,(function(e,t){return{"& > :not(style) + :not(style)":(0,s.Z)({margin:0},"margin".concat((i=t?d[t]:n.direction,{row:"Left","row-reverse":"Right",column:"Top","column-reverse":"Bottom"}[i])),(0,l.NA)(r,e))};var i})))}return i})),Z=r.forwardRef((function(e,n){var t=(0,f.Z)({props:e,name:"MuiStack"}),s=(0,c.Z)(t),r=s.component,a=void 0===r?"div":r,l=s.direction,u=void 0===l?"column":l,d=s.spacing,Z=void 0===d?0:d,v=s.divider,g=s.children,j=(0,i.Z)(s,p),b={direction:u,spacing:Z};return(0,x.jsx)(m,(0,o.Z)({as:a,ownerState:b,ref:n},j,{children:v?h(g,v):g}))}));n.Z=Z}}]);
//# sourceMappingURL=3328.400f6df1.chunk.js.map